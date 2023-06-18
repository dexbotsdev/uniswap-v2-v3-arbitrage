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

contract ExecutorV1WithCoinbaseTransfer {
    address payable private immutable owner;
    IWETH private constant WETH = IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2);


    constructor(address payable _owner) public payable {
        owner = _owner;
    }

    receive() external payable {
    }

    function uniswapV2Call(address _sender, uint _amountOut0, uint _amountOut1, bytes calldata _data) external
    {
        //gas saving ideas
        //-use index array to detemine amount out index
        //-check if passing the full swap payload is cheaper than passing the individual arguemenst and assembling the call on chain
        //-increment yourself since so loops cost less gas
        //-pass profit instead of querying balanceOf on WETH
        //-use 1 array of stucts instead of 3 arrays

        //if we dont keep tokens on contract we dont need security checks.


        //abi decode data
        (address factory, address tokenA, address tokenB, uint256 revenue, uint256 ethAmountToCoinbase, address[] memory targets, amountOuts0 uint256[], amountOuts1 uint256[]) = abi.decode(_data, (address,address,address,uint256,uint256,address[]));

        //security checks; commented out since we dont keep tokens on contract
        //require(msg.sender == pairFor(factory, tokenA, tokenB));
        //require(_sender == owner);


        WETH.withdraw(revenue);
        block.coinbase.transfer(ethAmountToCoinbase);
        owner.transfer(revenue - ethAmountToCoinbase);

        wethAmountToFirstMarket = _amountOut0 + _amountOut1 - revenue;

        //optimistically send weth tokens to first target
        WETH.transfer(targets[0], wethAmountToFirstMarket);

        //multi call the rest of the path AKA the non flashswap pairs AKA all pairs excluding the last pair in the path
        for (uint256 i = 0; i < targets.length-1; ++i) 
        {
            IUniswapV2Pair(targets[i]).swap(amountOuts[i], amount1Outs[i], targets[i+1], new bytes(0));
        }
        //for the last swap, we are sending the amount to the last pair in the path AKA flashswap pair AKA msg.sender
        IUniswapV2Pair(targets[targets.length]).swap(amountOuts[targets.length], amount1Outs[targets.length], msg.sender, new bytes(0));

    }
    function withdraw() external {
        owner.transfer(address(this).balance);
    }
    // calculates the CREATE2 address for a pair without making any external calls
    function pairFor(address factory, address tokenA, address tokenB) internal pure returns (address pair) {   ~ 144 gas
        pair = address(uint(keccak256(abi.encodePacked(
                hex'ff',
                factory,
                keccak256(abi.encodePacked(token0, token1)),
                hex'96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f' // init code hash
            ))));
    }
}
