// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

interface IUniswapV2Pair {
    function token0() external view returns (address);
    function token1() external view returns (address);
    function swap(uint amount0Out, uint amount1Out, address to, bytes calldata data) external;
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
}

contract GetUniswapV2Reserves {
    constructor(
        address[] memory pools
    ) {

        //for each pool get the reserves
        uint112[] memory reserve0s = new uint112[](pools.length);
        uint112[] memory reserve1s = new uint112[](pools.length);

        for (uint256 i=0;i<pools.length;i++) {
            (uint112 reserve0, uint112 reserve1,) = IUniswapV2Pair(pools[i]).getReserves();
            reserve0s[i] = reserve0;
            reserve1s[i] = reserve1;
        }

        bytes memory data = abi.encode(block.number,reserve0s, reserve1s);
        assembly {
            return(add(data, 32), mload(data))
        }
    }
}
