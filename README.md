
# 🧠 Idea Overview

This is a simple demonstration of how cellular automata can be used for data encryption. It uses a one-dimensional cellular automaton, inspired by ideas first introduced by Stephen Wolfram in 1985.

More details in Wolfram's original paper:
- Cryptography with Cellular Automata ([PDF][https://content.wolfram.com/uploads/sites/34/2020/07/cryptography-cellular-automata.pdf])

## 🔧 How It Works

The program includes several helper functions to handle data transformations and visualization:

- stringToBytes: Converts a string into a byte slice.
- parseBinToHex: Converts a byte slice into a hexadecimal string.
- show: Visualizes a byte slice (printing for 0 and | for 1).
- f: The core automaton rule function, defined as:
    f(y1, y2, y3) = y1 ^ (y2 | y3)

This rule is applied across the data to simulate the cellular automaton's evolution.

## Example

We encrypt the string "Hello" by converting it into binary and applying the cellular automaton transformation step by step. Each step shows how the data evolves over time.


```
Hello
Binary presentation:
0100100001100101011011000110110001101111
 |  |    ||  | | || ||   || ||   || ||||
 |||||  || ||| | |  | | ||  | | ||  |   
||    |||  |   | |||| | | ||| | | ||||  
| |  ||  |||| || |    | | |   | | |   ||
  |||| |||    |  ||  || | || || | || || 
 ||    |  |  ||||| |||  | |  |  | |  | |
 | |  ||||||||     |  ||| ||||||| |||| |
 | ||||       |   |||||   |       |    |
 | |   |     ||| ||    | |||     |||  ||
 | || |||   ||   | |  || |  |   ||  ||| 
|| |  |  | || | || ||||  ||||| || |||  |
   ||||||| |  | |  |   |||     |  |  |||
| ||       |||| ||||| ||  |   ||||||||  
| | |     ||    |     | |||| ||       ||
  | ||   || |  |||   || |    | |     || 
 || | | ||  ||||  | ||  ||  || ||   || |
 |  | | | |||   ||| | ||| |||  | | ||  |
 |||| | | |  | ||   | |   |  ||| | | |||
 |    | | |||| | | || || |||||   | | |  
|||  || | |    | | |  |  |    | || | || 
|  |||  | ||  || | ||||||||  || |  | |  
||||  ||| | |||  | |       |||  |||| |||
    |||   | |  ||| ||     ||  |||    |  
   ||  | || ||||   | |   || |||  |  ||| 
  || ||| |  |   | || || ||  |  ||||||  |
|||  |   ||||| || |  |  | ||||||     |||
   |||| ||     |  ||||||| |     |   ||  
  ||    | |   |||||       ||   ||| || | 
 || |  || || ||    |     || | ||   |  ||
 |  ||||  |  | |  |||   ||  | | | ||||| 
|||||   |||||| ||||  | || ||| | | |    |
     | ||      |   ||| |  |   | | ||  ||
|   || | |    ||| ||   ||||| || | | ||| 
|| ||  | ||  ||   | | ||     |  | | |   
|  | ||| | ||| | || | | |   ||||| | || |
 ||| |   | |   | |  | | || ||     | |  |
 |   || || || || |||| | |  | |   || ||||
 || ||  |  |  |  |    | |||| || ||  |   
||  | |||||||||||||  || |    |  | ||||  
| ||| |            |||  ||  ||||| |   ||
  |   ||          ||  ||| |||     || || 
0010001100000000001100111011100000110110
Encode result: 230033B836
```
