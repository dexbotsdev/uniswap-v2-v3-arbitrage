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

contract GetUniswapV3TicksV2 {
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

            int16 wordPos;
            {
                //get current tick
                (,int24 tick, , , , , ) = poolInstance.slot0();

                //get populated ticks for wordPos, wordPos-1, wordPos+1
                int24 tickSpacing = poolInstance.tickSpacing();
                int24 compressedTick = tickSpacing;
                if (tick < 0 && tick % tickSpacing != 0) compressedTick--;
                
                wordPos = int16(compressedTick >> 8);
            }

            uint size = 0;
            int16 bottomWord = 0;
            int16 topWord =0;
            //traverse down untill poluatedTick.length = 0
            while (true) {
                PopulatedTick[] memory populatedTicks = TickLens(0xbfd8137f7d1516D3ea5cA83523914859ec47F573).getPopulatedTicksInWord(pools[i], wordPos);
                if (populatedTicks.length == 0) {
                    bottomWord= wordPos+1;
                    break;
                } else {
                    wordPos--;
                }
            }

            wordPos = bottomWord;

            //traverse up untill poluatedTick.length = 0 and record size
            while (true) {
                PopulatedTick[] memory populatedTicks = TickLens(0xbfd8137f7d1516D3ea5cA83523914859ec47F573).getPopulatedTicksInWord(pools[i], wordPos);
                if (populatedTicks.length == 0) {
                    topWord = wordPos-1;
                    break;
                } else {
                    size += populatedTicks.length;
                    wordPos++;
                }
            }

            //declare the sizes of the columns
            indexArr[i] = new int24[](size);
            liquidityNetArr[i] = new int128[](size);
            liquidityGrossArr[i] = new uint128[](size);

            wordPos = bottomWord;

            //traverse up untill poluatedTick.length = 0 and record size
            //for each word
            //for each tick within word

            uint currPos = 0;
            for (int j = bottomWord; j <= topWord; j++){
                PopulatedTick[] memory populatedTicks = TickLens(0xbfd8137f7d1516D3ea5cA83523914859ec47F573).getPopulatedTicksInWord(pools[i], wordPos);
                for (uint256 k = 0; k < populatedTicks.length; ++k) {
                    indexArr[i][k] = populatedTicks[k].tick;
                    liquidityNetArr[i][k] = populatedTicks[k].liquidityNet;
                    liquidityGrossArr[i][k] = populatedTicks[k].liquidityGross;
                }
                currPos += populatedTicks.length;
                wordPos++;
            }
        }  
        

        bytes memory data = abi.encode(block.number, indexArr, liquidityNetArr, liquidityGrossArr);
        assembly {
            return(add(data, 32), mload(data))
        }
    }
}
