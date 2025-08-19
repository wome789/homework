pragma solidity ^0.6.0;

import "../../GSN/Context.sol";
import "./IERC777.sol";
import "./IERC777Recipient.sol";
import "./IERC777Sender.sol";
import "../../token/ERC20/IERC20.sol";
import "../../math/SafeMath.sol";
import "../../utils/Address.sol";
import "../../introspection/IERC1820Registry.sol";
import "../../Initializable.sol";

/**
 * @dev {IERC777} 接口的实现。
 *
 * 此实现与代币创建方式无关。这意味着
 * 必须在派生合约中使用 {_mint} 添加供应机制。
 *
 * 按照EIP规范，此合约包含对ERC20的支持：与它交互时，
 * 可以安全地使用ERC777和ERC20接口。
 * 代币移动时会发出 {IERC777-Sent} 和 {IERC20-Transfer} 事件。
 *
 * 此外，{IERC777-granularity} 值硬编码为 `1`，这意味着
 * 在创建、移动或销毁的代币数量方面没有特殊限制。
 * 这使得与ERC20应用程序的集成变得无缝。
 */
contract ERC777UpgradeSafe is Initializable, ContextUpgradeSafe, IERC777, IERC20 {
    using SafeMath for uint256;
    using Address for address;

    // ERC1820注册表常量
    IERC1820Registry constant internal _ERC1820_REGISTRY = IERC1820Registry(0x1820a4B7618BdE71Dce8cdc73aAB6C95905faD24);

    // 地址到余额的映射
    mapping(address => uint256) private _balances;

    // 总供应量
    uint256 private _totalSupply;

    // 代币名称
    string private _name;
    // 代币符号
    string private _symbol;

    // 我们内联以下哈希的结果，因为Solidity在编译时无法解析它们。
    // 参见 https://github.com/ethereum/solidity/issues/4024。

    // keccak256("ERC777TokensSender")
    bytes32 constant private _TOKENS_SENDER_INTERFACE_HASH =
        0x29ddb589b1fb5fc7cf394961c1adf5f8c6454761adf795e67fe149f658abe895;

    // keccak256("ERC777TokensRecipient")
    bytes32 constant private _TOKENS_RECIPIENT_INTERFACE_HASH =
        0xb281fc8c12954d22544db45de3159a39272895b169a852b314f9cc762e44c53b;

    // 这从未被读取 - 仅用于响应defaultOperators查询。
    address[] private _defaultOperatorsArray;

    // 不可变的，但账户可以撤销它们（在__revokedDefaultOperators中跟踪）。
    mapping(address => bool) private _defaultOperators;

    // 对于每个账户，其操作员和已撤销的默认操作员的映射。
    mapping(address => mapping(address => bool)) private _operators;
    mapping(address => mapping(address => bool)) private _revokedDefaultOperators;

    // ERC20-授权
    mapping (address => mapping (address => uint256)) private _allowances;

    /**
     * @dev `defaultOperators` 可能是一个空数组。
     */

    function __ERC777_init(
        string memory name,
        string memory symbol,
        address[] memory defaultOperators
    ) internal initializer {
        __Context_init_unchained();
        __ERC777_init_unchained(name, symbol, defaultOperators);
    }

    function __ERC777_init_unchained(
        string memory name,
        string memory symbol,
        address[] memory defaultOperators
    ) internal initializer {

        _name = name;
        _symbol = symbol;

        _defaultOperatorsArray = defaultOperators;
        for (uint256 i = 0; i < _defaultOperatorsArray.length; i++) {
            _defaultOperators[_defaultOperatorsArray[i]] = true;
        }

        // 注册接口
        _ERC1820_REGISTRY.setInterfaceImplementer(address(this), keccak256("ERC777Token"), address(this));
        _ERC1820_REGISTRY.setInterfaceImplementer(address(this), keccak256("ERC20Token"), address(this));

    }

    /**
     * @dev 参见 {IERC777-name}。
     */
    function name() public view override returns (string memory) {
        return _name;
    }

    /**
     * @dev 参见 {IERC777-symbol}。
     */
    function symbol() public view override returns (string memory) {
        return _symbol;
    }

    /**
     * @dev 参见 {ERC20-decimals}。
     *
     * 根据 [ERC777 EIP](https://eips.ethereum.org/EIPS/eip-777#backward-compatibility)，
     * 始终返回18。
     */
    function decimals() public pure returns (uint8) {
        return 18;
    }

    /**
     * @dev 参见 {IERC777-granularity}。
     *
     * 此实现始终返回 `1`。
     */
    function granularity() public view override returns (uint256) {
        return 1;
    }

    /**
     * @dev 参见 {IERC777-totalSupply}。
     */
    function totalSupply() public view override(IERC20, IERC777) returns (uint256) {
        return _totalSupply;
    }

    /**
     * @dev 返回账户（`tokenHolder`）拥有的代币数量。
     */
    function balanceOf(address tokenHolder) public view override(IERC20, IERC777) returns (uint256) {
        return _balances[tokenHolder];
    }

    /**
     * @dev 参见 {IERC777-send}。
     *
     * 同时为ERC20兼容性发出 {IERC20-Transfer} 事件。
     */
    function send(address recipient, uint256 amount, bytes memory data) public override  {
        _send(_msgSender(), recipient, amount, data, "", true);
    }

    /**
     * @dev 参见 {IERC20-transfer}。
     *
     * 与 `send` 不同，如果 `recipient` 是合约，则不需要实现 {IERC777Recipient} 接口。
     *
     * 同时发出 {Sent} 事件。
     */
    function transfer(address recipient, uint256 amount) public override returns (bool) {
        require(recipient != address(0), "ERC777: 转移到零地址");

        address from = _msgSender();

        _callTokensToSend(from, from, recipient, amount, "", "");

        _move(from, from, recipient, amount, "", "");

        _callTokensReceived(from, from, recipient, amount, "", "", false);

        return true;
    }

    /**
     * @dev 参见 {IERC777-burn}。
     *
     * 同时为ERC20兼容性发出 {IERC20-Transfer} 事件。
     */
    function burn(uint256 amount, bytes memory data) public override  {
        _burn(_msgSender(), amount, data, "");
    }

    /**
     * @dev 参见 {IERC777-isOperatorFor}。
     */
    function isOperatorFor(
        address operator,
        address tokenHolder
    ) public view override returns (bool) {
        return operator == tokenHolder ||
            (_defaultOperators[operator] && !_revokedDefaultOperators[tokenHolder][operator]) ||
            _operators[tokenHolder][operator];
    }

    /**
     * @dev 参见 {IERC777-authorizeOperator}。
     */
    function authorizeOperator(address operator) public override  {
        require(_msgSender() != operator, "ERC777: 将自己授权为操作员");

        if (_defaultOperators[operator]) {
            delete _revokedDefaultOperators[_msgSender()][operator];
        } else {
            _operators[_msgSender()][operator] = true;
        }

        emit AuthorizedOperator(operator, _msgSender());
    }

    /**
     * @dev 参见 {IERC777-revokeOperator}。
     */
    function revokeOperator(address operator) public override  {
        require(operator != _msgSender(), "ERC777: 撤销自己作为操作员");

        if (_defaultOperators[operator]) {
            _revokedDefaultOperators[_msgSender()][operator] = true;
        } else {
            delete _operators[_msgSender()][operator];
        }

        emit RevokedOperator(operator, _msgSender());
    }

    /**
     * @dev 参见 {IERC777-defaultOperators}。
     */
    function defaultOperators() public view override returns (address[] memory) {
        return _defaultOperatorsArray;
    }

    /**
     * @dev 参见 {IERC777-operatorSend}。
     *
     * 发出 {Sent} 和 {IERC20-Transfer} 事件。
     */
    function operatorSend(
        address sender,
        address recipient,
        uint256 amount,
        bytes memory data,
        bytes memory operatorData
    )
    public override
    {
        require(isOperatorFor(_msgSender(), sender), "ERC777: 调用者不是持有者的操作员");
        _send(sender, recipient, amount, data, operatorData, true);
    }

    /**
     * @dev 参见 {IERC777-operatorBurn}。
     *
     * 发出 {Burned} 和 {IERC20-Transfer} 事件。
     */
    function operatorBurn(address account, uint256 amount, bytes memory data, bytes memory operatorData) public override {
        require(isOperatorFor(_msgSender(), account), "ERC777: 调用者不是持有者的操作员");
        _burn(account, amount, data, operatorData);
    }

    /**
     * @dev 参见 {IERC20-allowance}。
     *
     * 注意，操作员和授权概念是正交的：操作员可能没有授权，
     * 而有授权的账户本身可能不是操作员。
     */
    function allowance(address holder, address spender) public view override returns (uint256) {
        return _allowances[holder][spender];
    }

    /**
     * @dev 参见 {IERC20-approve}。
     *
     * 注意，账户不能由其操作员发出授权。
     */
    function approve(address spender, uint256 value) public override returns (bool) {
        address holder = _msgSender();
        _approve(holder, spender, value);
        return true;
    }

   /**
    * @dev 参见 {IERC20-transferFrom}。
    *
    * 注意，操作员和授权概念是正交的：操作员不能调用 `transferFrom`
    * （除非他们有授权），而有授权的账户不能调用 `operatorSend`
    * （除非他们是操作员）。
    *
    * 发出 {Sent}、{IERC20-Transfer} 和 {IERC20-Approval} 事件。
    */
    function transferFrom(address holder, address recipient, uint256 amount) public override returns (bool) {
        require(recipient != address(0), "ERC777: 转移到零地址");
        require(holder != address(0), "ERC777: 从零地址转移");

        address spender = _msgSender();

        _callTokensToSend(spender, holder, recipient, amount, "", "");

        _move(spender, holder, recipient, amount, "", "");
        _approve(holder, spender, _allowances[holder][spender].sub(amount, "ERC777: 转移金额超过授权"));

        _callTokensReceived(spender, holder, recipient, amount, "", "", false);

        return true;
    }

    /**
     * @dev 创建 `amount` 个代币并将它们分配给 `account`，增加总供应量。
     *
     * 如果为 `account` 注册了发送钩子，将使用 `operator`、`data` 和 `operatorData`
     * 调用相应的函数。
     *
     * 参见 {IERC777Sender} 和 {IERC777Recipient}。
     *
     * 发出 {Minted} 和 {IERC20-Transfer} 事件。
     *
     * 要求
     *
     * - `account` 不能是零地址。
     * - 如果 `account` 是合约，它必须实现 {IERC777Recipient} 接口。
     */
    function _mint(
        address account,
        uint256 amount,
        bytes memory userData,
        bytes memory operatorData
    )
    internal virtual
    {
        require(account != address(0), "ERC777: 铸造到零地址");

        address operator = _msgSender();

        _beforeTokenTransfer(operator, address(0), account, amount);

        // 更新状态变量
        _totalSupply = _totalSupply.add(amount);
        _balances[account] = _balances[account].add(amount);

        _callTokensReceived(operator, address(0), account, amount, userData, operatorData, true);

        emit Minted(operator, account, amount, userData, operatorData);
        emit Transfer(address(0), account, amount);
    }

    /**
     * @dev 发送代币
     * @param from 代币持有者地址
     * @param to 接收者地址
     * @param amount 要转移的代币数量
     * @param userData 代币持有者提供的额外信息（如果有）
     * @param operatorData 操作员提供的额外信息（如果有）
     * @param requireReceptionAck 如果为true，合约接收者需要实现ERC777TokensRecipient
     */
    function _send(
        address from,
        address to,
        uint256 amount,
        bytes memory userData,
        bytes memory operatorData,
        bool requireReceptionAck
    )
        internal
    {
        require(from != address(0), "ERC777: 从零地址发送");
        require(to != address(0), "ERC777: 发送到零地址");

        address operator = _msgSender();

        _callTokensToSend(operator, from, to, amount, userData, operatorData);

        _move(operator, from, to, amount, userData, operatorData);

        _callTokensReceived(operator, from, to, amount, userData, operatorData, requireReceptionAck);
    }

    /**
     * @dev 销毁代币
     * @param from 代币持有者地址
     * @param amount 要销毁的代币数量
     * @param data 代币持有者提供的额外信息
     * @param operatorData 操作员提供的额外信息（如果有）
     */
    function _burn(
        address from,
        uint256 amount,
        bytes memory data,
        bytes memory operatorData
    )
        internal virtual
    {
        require(from != address(0), "ERC777: 从零地址销毁");

        address operator = _msgSender();

        _beforeTokenTransfer(operator, from, address(0), amount);

        _callTokensToSend(operator, from, address(0), amount, data, operatorData);

        // 更新状态变量
        _balances[from] = _balances[from].sub(amount, "ERC777: 销毁金额超过余额");
        _totalSupply = _totalSupply.sub(amount);

        emit Burned(operator, from, amount, data, operatorData);
        emit Transfer(from, address(0), amount);
    }

    function _move(
        address operator,
        address from,
        address to,
        uint256 amount,
        bytes memory userData,
        bytes memory operatorData
    )
        private
    {
        _beforeTokenTransfer(operator, from, to, amount);

        _balances[from] = _balances[from].sub(amount, "ERC777: 转移金额超过余额");
        _balances[to] = _balances[to].add(amount);

        emit Sent(operator, from, to, amount, userData, operatorData);
        emit Transfer(from, to, amount);
    }

    function _approve(address holder, address spender, uint256 value) internal {
        // TODO: 如果此函数变为内部函数或在新的调用点调用，则恢复此require语句。目前是不必要的。
        //require(holder != address(0), "ERC777: 从零地址授权");
        require(spender != address(0), "ERC777: 授权到零地址");

        _allowances[holder][spender] = value;
        emit Approval(holder, spender, value);
    }

    /**
     * @dev 如果接口已注册，则调用 from.tokensToSend()
     * @param operator 请求转移的操作员地址
     * @param from 代币持有者地址
     * @param to 接收者地址
     * @param amount 要转移的代币数量
     * @param userData 代币持有者提供的额外信息（如果有）
     * @param operatorData 操作员提供的额外信息（如果有）
     */
    function _callTokensToSend(
        address operator,
        address from,
        address to,
        uint256 amount,
        bytes memory userData,
        bytes memory operatorData
    )
        private
    {
        address implementer = _ERC1820_REGISTRY.getInterfaceImplementer(from, _TOKENS_SENDER_INTERFACE_HASH);
        if (implementer != address(0)) {
            IERC777Sender(implementer).tokensToSend(operator, from, to, amount, userData, operatorData);
        }
    }

    /**
     * @dev 如果接口已注册，则调用 to.tokensReceived()。如果接收者是合约但
     * 没有为接收者注册tokensReceived()，则回滚
     * @param operator 请求转移的操作员地址
     * @param from 代币持有者地址
     * @param to 接收者地址
     * @param amount 要转移的代币数量
     * @param userData 代币持有者提供的额外信息（如果有）
     * @param operatorData 操作员提供的额外信息（如果有）
     * @param requireReceptionAck 如果为true，合约接收者需要实现ERC777TokensRecipient
     */
    function _callTokensReceived(
        address operator,
        address from,
        address to,
        uint256 amount,
        bytes memory userData,
        bytes memory operatorData,
        bool requireReceptionAck
    )
        private
    {
        address implementer = _ERC1820_REGISTRY.getInterfaceImplementer(to, _TOKENS_RECIPIENT_INTERFACE_HASH);
        if (implementer != address(0)) {
            IERC777Recipient(implementer).tokensReceived(operator, from, to, amount, userData, operatorData);
        } else if (requireReceptionAck) {
            require(!to.isContract(), "ERC777: 代币接收者合约没有实现ERC777TokensRecipient");
        }
    }

    /**
     * @dev 在任何代币转移之前调用的钩子。这包括
     * 对 {send}、{transfer}、{operatorSend}、铸造和销毁的调用。
     *
     * 调用条件：
     *
     * - 当 `from` 和 `to` 都不为零时，``from`` 的 `tokenId` 将
     * 转移给 `to`。
     * - 当 `from` 为零时，将为 `to` 铸造 `tokenId`。
     * - 当 `to` 为零时，``from`` 的 `tokenId` 将被销毁。
     * - `from` 和 `to` 永远不会都为零。
     *
     * 要了解有关钩子的更多信息，请参阅 xref:ROOT:extending-contracts.adoc#using-hooks[使用钩子]。
     */
    function _beforeTokenTransfer(address operator, address from, address to, uint256 tokenId) internal virtual { }

    uint256[41] private __gap;
}
