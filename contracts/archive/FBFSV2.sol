//SPDX-License-Identifier: UNLICENSED
pragma solidity 0.6.12;

pragma experimental ABIEncoderV2;

interface IERC20 {
    event Approval(address indexed owner, address indexed spender, uint value);
    event Transfer(address indexed from, address indexed to, uint value);

    function name() external view returns (string memory);
    function symbol() external view returns (string memory);
    function decimals() external view returns (uint8);
    function totalSupply() external view returns (uint);
    function balanceOf(address owner) external view returns (uint);
    function allowance(address owner, address spender) external view returns (uint);

    function approve(address spender, uint value) external returns (bool);
    function transfer(address to, uint value) external returns (bool);
    function transferFrom(address from, address to, uint value) external returns (bool);
}

interface IWETH is IERC20 {
    function deposit() external payable;
    function withdraw(uint) external;
}

interface IUniswapV2Pair {
    function token0() external view returns (address);
    function token1() external view returns (address);
    function swap(uint amount0Out, uint amount1Out, address to, bytes calldata data) external;
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
}
// This contract simply calls multiple targets sequentially, ensuring WETH balance before and after

contract FBFS {
    address payable private immutable owner;
    IWETH private constant WETH = IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2);


    constructor(address payable _owner) public payable {
        owner = _owner;
    }

    receive() external payable {
    }

    function uniswapV2Call(address _sender, uint _amount0, uint _amount1, bytes calldata _data) external
    {
        address token0 = IUniswapV2Pair(msg.sender).token0(); // fetch the address of token0
        address token1 = IUniswapV2Pair(msg.sender).token1(); // fetch the address of token1
        //require(msg.sender == IUniswapV2Factory(0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f).getPair(token0, token1));
        require(_sender == owner);
        //abi decode data
        (uint256 wethAmountToFirstMarket,uint256 _ethAmountToCoinbase,address[] memory targets, bytes[] memory payloads) = abi.decode(_data, (uint256,uint256,address[],bytes[]));

        //optimistically send tokens to first target
        WETH.transfer(targets[0], wethAmountToFirstMarket);

        //multi call the rest of the path
        for (uint256 i = 0; i < targets.length; i++) 
        {
            targets[i].call(payloads[i]);
        }

        //WETH.withdraw(WETH.balanceOf(address(this)));

        //owner.transfer(address(this).balance); // our win
    }
}
