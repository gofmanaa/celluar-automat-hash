This is a simple example of how cellular automata can be used to encrypt data. An example uses a one-dimensional cellular automaton.

## How it works

The code includes several helper functions to assist with these operations:

* `stringToBytes` converts a string to a slice of bytes

* `parseBinToHex` converts a slice of bytes to a hexadecimal string

* `show` prints a slice of bytes as a visual representation (using 
` ` for 0 and `|` for 1)
* `f` is the custom algorithm that calculates the new slice of bytes. It takes three arguments, y1, y2, and y3, and returns the result of `y1 ^ (y2 | y3)`.

More detail in the link below
https://content.wolfram.com/uploads/sites/34/2020/07/cryptography-cellular-automata.pdf