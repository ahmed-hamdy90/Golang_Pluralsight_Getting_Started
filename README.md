#### This Repository Follow Pluralsight course for [GO: Getting Started](https://app.pluralsight.com/library/courses/getting-started-with-go)

##### Examples:
 every clip has include code or demo has been written on separated folder under folder for which Chapter this clip exists, Here List of Chapters and clips

- [Introduction](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch02)
	- [Demo: Hello World](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch02/05)
- [Starting a Project](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch03)
  - [Creating a Project](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch03/11)
- [Working with Primitive Data Types](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch04)
	- [Declaring Variables with Primitive Data Types](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch03/14)
	- [Working with Pointers](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch03/15)
	- [Creating Constants](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch03/16)
	- [Using Iota and Constant Expressions](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch03/17)
- [Working with Collections](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch05)
  - [Creating Arrays](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch05/20)
  - [Working with Slices](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch05/21)
  - [Using the Map Data Type](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch05/22)
  - [Working with Structs](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch05/23)
  - [Demo: Adding Variables to the Webservice](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch05/24)
- [Creating Functions and Methods](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch06)
  - [Creating Functions](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch06/26)
  - [Adding Parameters to Functions](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch06/27)
  - [Returning Data from Functions](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch06/28)
  - [Using Methods to Add Behaviors to a Type](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch06/29)
  - [Demo: Adding Functions to the Webservice](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch06/30)
  - [Demo: Creating Methods](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch06/31)
  - [Demo: Implementing Interfaces](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch06/32)
  - [Demo: Starting the Webservice](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch06/33)
- [Controlling Program Flow](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch07)
  - [Creating Loops that Terminate Based on a Condition](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch07/37)
  - [Using Conditional Loops with Post Clauses](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch07/38)
  - [Creating Infinite Loops](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch07/39)
  - [Looping Over Collections](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch07/40)
  - [Working with the Panic Function](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch07/41)
  - [Creating If Statements](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch07/42)
  - [Demo: Adding Loops and Branches to the WebService](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch07/43)
  - [Writing Switch Statements](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch07/44)
  - [Demo: Adding Switch Statements to the WebService](https://github.com/ahmed-hamdy90/Golang_Pluralsight_Getting_Started/tree/master/Ch07/45)

##### Run Examples:

 You can Run every example using Docker(every one has his own Dockerfile file), if you don't have Docker , you can install Docker Engine from [Here](https://docs.docker.com/engine/install)

 - First build docker image
```shell
docker build -t golangtestapp .
```

 - Second run built image
```shell
docker run --name golangtestapp -p 3000:3000 golangtestapp
```
