// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;


interface IUniswapV2Facotry {
    function getPair(address tokenA, address tokenB) external view returns (address pair);
    function allPairs(uint) external view returns (address pair);
    function allPairsLength() external view returns (uint);
}
struct PopulatedTick {
    int24 tick;
    int128 liquidityNet;
    uint128 liquidityGross;
}
interface TickLens{

    function getPopulatedTicksInWord(address pool, int16 tickBitmapIndex) external view returns (PopulatedTick[] memory populatedTicks);
}

interface IUniswapV3Pool{
    function token0() external view returns (address);
    function token1() external view returns (address);
    function slot0() external view returns (uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked);
    function fee() external view returns (uint24);
    function tickSpacing() external view returns (int24);
    function observationCardinalityNext() external view returns (uint16);
    function observationCardinality() external view returns (uint16);
    function observationIndex() external view returns (uint16);
    function observation(uint256 index) external view returns (uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized);
    function liquidity() external view returns (uint128);
    function sqrtPriceX96() external view returns (uint160);
    function tickCumulative(uint256 index) external view returns (int56);
    function secondsPerLiquidityCumulativeX128(uint256 index) external view returns (uint160);
    function tick() external view returns (int24);
    function maxLiquidityPerTick() external view returns (uint128);
    function liquidityCumulative(uint256 index) external view returns (uint256);
    function initialize(uint160 sqrtPriceX96) external;
    function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes calldata data) external returns (uint256 amount0, uint256 amount1);
    function burn(int24 tickLower, int24 tickUpper, uint128 amount) external returns (uint256 amount0, uint256 amount1);
    function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes calldata data) external returns (int256 amount0, int256 amount1);
    function flash(address recipient, uint256 amount0, uint256 amount1, bytes calldata data) external;
    function increaseObservationCardinalityNext(uint16 observationCardinalityNext) external;
    function decreaseObservationCardinalityNext(uint16 observationCardinalityNext) external;
    function initialize(uint160 sqrtPriceX96, uint16 _observationCardinalityLimit, uint16 _observationCardinalityNext) external;
    function collectProtocol(address token0, address token1, uint128 amount0Requested, uint128 amount1Requested) external returns (uint128 amount0, uint128 amount1);
}


//ticklens address

contract GetUniswapV3Ticks {
      // Define the event with a bytes parameter
    event BytesEmitted(bytes data);
    constructor(
        address[] memory pools 
    ){

      //console.log(pools.length);

        //  Token0       common.Address
        //Token1       common.Address
        //SqrtRatioX96 *big.Int
        //Liquidity    *big.Int
        //TickCurrent  int

        //TickInfoMap map[int]*Tick


        uint len = pools.length;

        int24[][] memory indexArr = new int24[][](len);
        int128[][] memory liquidityNetArr = new int128[][](len);
        uint128[][] memory liquidityGrossArr = new uint128[][](len);
        

        for (uint256 i = 0; i < len; ++i) {
            IUniswapV3Pool poolInstance = IUniswapV3Pool(pools[i]);

            //get current tick
            (,int24 tick, , , , , ) = poolInstance.slot0();

            //get populated ticks for wordPos, wordPos-1, wordPos+1
            int24 tickSpacing = poolInstance.tickSpacing();
            int24 compressedTick = tickSpacing;
            if (tick < 0 && tick % tickSpacing != 0) compressedTick--;
             
            int16 wordPos = int16(compressedTick >> 8);
            PopulatedTick[] memory populatedTicksMinus1 = TickLens(0xbfd8137f7d1516D3ea5cA83523914859ec47F573).getPopulatedTicksInWord(pools[i], wordPos-1);
            PopulatedTick[] memory populatedTicks = TickLens(0xbfd8137f7d1516D3ea5cA83523914859ec47F573).getPopulatedTicksInWord(pools[i], wordPos);
            PopulatedTick[] memory populatedTicksPlus1 = TickLens(0xbfd8137f7d1516D3ea5cA83523914859ec47F573).getPopulatedTicksInWord(pools[i], wordPos+1);

            //combine all the populated ticks
            // PopulatedTick[] memory allPopulatedTicks;

            // for (uint256 j = 0; j < populatedTicksMinus1.length+populatedTicks.length+populatedTicksPlus1.length; ++j) {
            //     allPopulatedTicks[j] = populatedTicksMinus1[j];
            // }

            //declare the sizes of the columns
            indexArr[i] = new int24[](populatedTicksMinus1.length + populatedTicks.length + populatedTicksPlus1.length);
            liquidityNetArr[i] = new int128[](populatedTicksMinus1.length + populatedTicks.length + populatedTicksPlus1.length);
            liquidityGrossArr[i] = new uint128[](populatedTicksMinus1.length + populatedTicks.length + populatedTicksPlus1.length);

            //indexArr[i] = new int24[](populatedTicks.length);
            // liquidityNetArr[i] = new int128[](populatedTicks.length);
            // liquidityGrossArr[i] = new uint128[](populatedTicks.length);

        // console.log("populatedTicksMinus1.length", populatedTicksMinus1.length);
        // console.log("populatedTicks.length", populatedTicks.length);
        // console.log("populatedTicksPlus1.length", populatedTicksPlus1.length);

        // console.log("indexArr[i].length", indexArr[i].length);
        // console.log("liquidityNetArr[i].length", liquidityNetArr[i].length);
        // console.log("liquidityGrossArr[i].length", liquidityGrossArr[i].length);

            uint256 currPos =0;

            for (uint256 j = currPos; j < populatedTicksMinus1.length; ++j) {
                indexArr[i][j] = populatedTicksMinus1[j].tick;
                liquidityNetArr[i][j] = populatedTicksMinus1[j].liquidityNet;
                liquidityGrossArr[i][j] = populatedTicksMinus1[j].liquidityGross;
            }

            for (uint256 j = currPos; j < currPos+populatedTicks.length; ++j) {
                indexArr[i][j] = populatedTicks[j-currPos].tick;
                liquidityNetArr[i][j] = populatedTicks[j-currPos].liquidityNet;
                liquidityGrossArr[i][j] = populatedTicks[j-currPos].liquidityGross;
            }

            for (uint256 j = currPos; j < currPos + populatedTicksPlus1.length; ++j) {
                indexArr[i][j] = populatedTicksPlus1[j-currPos].tick;
                liquidityNetArr[i][j] = populatedTicksPlus1[j-currPos].liquidityNet;
                liquidityGrossArr[i][j] = populatedTicksPlus1[j-currPos].liquidityGross;
            }

            // for (uint256 j = 0; j < populatedTicks.length; ++j) {
            //     indexArr[i][j] = populatedTicks[j].tick;
            //     liquidityNetArr[i][j] = populatedTicks[j].liquidityNet;
            //     liquidityGrossArr[i][j] = populatedTicks[j].liquidityGross;
            // }
        }  

        bytes memory data = abi.encode(block.number, indexArr, liquidityNetArr, liquidityGrossArr);
        assembly {
            return(add(data, 32), mload(data))
        }
    }
}
