// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

interface IUniswapV2Facotry {
    function getPair(address tokenA, address tokenB) external view returns (address pair);
    function allPairs(uint) external view returns (address pair);
    function allPairsLength() external view returns (uint);
}

interface IUniswapV2Pair {
    function token0() external view returns (address);
    function token1() external view returns (address);
    function swap(uint amount0Out, uint amount1Out, address to, bytes calldata data) external;
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
}

contract GetUniswapV2Pools {
    constructor(
        address factory,
        uint start,
        uint end
    ) {

        address[] memory pairs = new address[](end-start);
        address[] memory token0s = new address[](end-start);
        address[] memory token1s = new address[](end-start);
        uint112[] memory reserve0s = new uint112[](end-start);
        uint112[] memory reserve1s = new uint112[](end-start);

        for (uint256 i = start; i < end; ++i) {
            address pair = IUniswapV2Facotry(factory).allPairs(i);
            pairs[i-start] = pair;
            token0s[i-start] = IUniswapV2Pair(pair).token0();
            token1s[i-start] = IUniswapV2Pair(pair).token1();
            (uint112 reserve0, uint112 reserve1,) = IUniswapV2Pair(pair).getReserves();
            reserve0s[i-start] = reserve0;
            reserve1s[i-start] = reserve1;

        }  


        bytes memory data = abi.encode(block.number, pairs, token0s, token1s, reserve0s, reserve1s);
        assembly {
            return(add(data, 32), mload(data))
        }
    }
}
