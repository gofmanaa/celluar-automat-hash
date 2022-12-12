This is a simple example of how cellular automata can be used to encrypt data. An example uses a one-dimensional cellular automaton.

I had an idea to write some code that uses a cellular automaton algorithm to encrypt data. It turned out that this idea was described by Stefan Wolfram back in 1985. For more detail in the link below:

https://content.wolfram.com/uploads/sites/34/2020/07/cryptography-cellular-automata.pdf

## How it works

The code includes several helper functions to assist with these operations:

* `stringToBytes` converts a string to a slice of bytes

* `parseBinToHex` converts a slice of bytes to a hexadecimal string

* `show` prints a slice of bytes as a visual representation (using 
` ` for 0 and `|` for 1)
* `f` is the custom algorithm that calculates the new slice of bytes. It takes three arguments, y1, y2, and y3, and returns the result of `y1 ^ (y2 | y3)`.

## Exsample:

We encrypt the string "Hello" and apply the algorithm of a one-dimensional cellular automaton to the binary representation of the string.

```
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
Encode result: 230033B836, [100011 0 110011 10111000 110110]
```