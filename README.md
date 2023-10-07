# Uniswap Arbitrage Bot 🤖

## About the Bot

- 🚧 This bot is still in the development phase.
- 🔄 Retrieves pools from both Uniswap v2 and v3.
- 🛡 Filters out tokens with low liquidity and questionable quality.
- 🔍 Implements Depth-First Search (DFS) to explore all possible paths.
- 🌟 Utilizes the Golden State Search method to determine the optimum input amount.
- 📊 Calculates potential revenue accurately based on Uniswap v2 & v3 math.
- ⚡ Incorporates flashswaps, so there's no capital needed upfront.
- 💼 Executions are managed through flashbots.
- 🚀 Features a custom smart contract with inline assembly to optimize for gas.
- 🌐 Uses a proprietary compression technique to further optimize gas usage.

## Building and Running the Bot

There are three commands to interact with the bot:

1. **Just Build**: 
```
python ./bot-builder.py build
```

2. **Build & Run**: 

```
python ./bot-builder.py buildRun
```

3. **Build & Run (Debug Mode)**: 

```
python ./bot-builder.py buildDebugRun
```

---

Contributions, feedback, and issues are welcome! Please ensure to follow our contribution guidelines and code of conduct.

---

Note: Ensure you have all the dependencies installed and have set up the environment variables correctly before building and running the bot.
