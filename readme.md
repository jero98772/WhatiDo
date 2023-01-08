# WhatIDo
What i do is an algorithm is for chose a final product with the materials given by input, the algoritm represent materials as graph

### Develompment continue
** TODO'S **

 - add numbers of elements needed to make something
 
# Explanation
'What i do' is an algorithm is for chose a final product with the materials given by input, the algoritm represent materials as graph ,this algorim work on directional graph to connect diferentes materials with their products, think as a function that recive a graph and list of items do you have

### Why is directional graph?

because we can make a house with concret and bricks but you don´t can make brick and concret with a house  

### Nodes

are materials and direccional

### Edges

are the things you can make

### programing notes

i use adyacent list to representate the graph because is more eficent than, but you can use matrix the algoritm work very simililar in a matrix (see <a href="">algoritm with matrix</a>)

### How it works for select a product?

to select the product that fits most of your available materials, search in the graph your items and put in a map (or python diccionary) the amout of ocurrenes of a products you can make,finalized with a map of the form {product:num of materials do you have} ending looking for the greatest occurrence and selecting the product with which it has it

### features 

- use adyacent list,(is optimized)

- web interface for play with algoritm

### Screenshots
![1](https://github.com/jero98772/WhatiDo/blob/main/misc/screenshots/1.png?raw=true)

![2](https://github.com/jero98772/WhatiDo/blob/main/misc/screenshots/2.png?raw=true)

### Installing
**Download repositories**

    git clone https://github.com/jero98772/whatIDo.git

**Run:**  

    go run .

### thanks to

 - Ricardo Mota Sohm for idea, (but i development)

 - Mauricio Toro for advices