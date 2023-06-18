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
    address private immutable owner;
    IWETH private constant WETH = IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2);


    constructor(address _owner) public payable {
        owner = _owner;
    }

    receive() external payable {
    }

    function flashSwapWeth(uint256 _ethAmountToCoinbase, address _flashSwapPair, bytes memory _flashSwapPayload) external payable {
    
        (bool _success, bytes memory _response) =_flashSwapPair.call(_flashSwapPayload);
        require(_success); _response;


        //after swaps are done, check if profitable

        require(WETH.balanceOf(address(this)) > _ethAmountToCoinbase);

        uint256 _ethBalance = address(this).balance;
        if (_ethBalance < _ethAmountToCoinbase) {
            WETH.withdraw(_ethAmountToCoinbase - _ethBalance);
        }

        WETH.transfer(owner, WETH.balanceOf(address(this))); // our win

        block.coinbase.transfer(_ethAmountToCoinbase);
    }

    function uniswapV2Call(address _sender, uint _amount0, uint _amount1, bytes calldata _data) external
    {
        assert(_sender == owner); // ensure that sender is this contract

        //abi decode data
        (uint256 wethAmountToFirstMarket,address[] memory targets, bytes[] memory payloads) = abi.decode(_data, (uint256,address[],bytes[]));

        //optimistically send tokens to first target
        WETH.transfer(targets[0], wethAmountToFirstMarket);

        //multi call the rest of the path
        for (uint256 i = 0; i < targets.length; i++) {

            (bool _success, bytes memory _response) = targets[i].call(payloads[i]);
            require(_success); _response;
        }
    }
}
