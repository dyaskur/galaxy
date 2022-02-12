# GALAXY MERCHANT TRADING GUIDE

### Problem Description

You decided to give up on earth after the latest financial collapse left 99.99% of the earth's population with 0.01% of the wealth. Luckily, with the scant sum of money that is left in your account, you are able to afford to rent a spaceship, leave earth, and fly all over the galaxy to sell common metals and dirt (which apparently is worth a lot). Buying and selling over the galaxy requires you to convert numbers and units, and you decided to write a program to help you.The numbers used for intergalactic transactions follows similar convention to the roman numerals and you have painstakingly collected the appropriate translation between them. Roman numerals are based on seven symbols:

Symbol Value:
I 1  
V 5  
X 10
L 50  
C 100  
D 500
M 1,000

Numbers are formed by combining symbols together and adding the values. For example, MMVI is 1000 + 1000 + 5 + 1 = 2006. Generally, symbols are placed in order of value, starting with the largest values. When smaller values precede larger values, the smaller values are subtracted from the larger values, and the result is added to the total. For example MCMXLIV = 1000 + (1000 − 100) + (50 − 10) + (5 − 1) = 1944.

The symbols "I", "X", "C", and "M" can be repeated three times in succession, but no more. (They may appear four times if the third and fourth are separated by a smaller value, such as XXXIX.) "D", "L", and "V" can never be repeated.

"I" can be subtracted from "V" and "X" only. "X" can be subtracted from "L" and "C" only. "C" can be subtracted from "D" and "M" only. "V", "L", and "D" can never be subtracted.

Only one small-value symbol may be subtracted from any large-value symbol.

A number written in Arabic numerals can be broken into digits. For example, 1903 is composed of 1, 9, 0, and 3. To write the Roman numeral, each of the non-zero digits should be treated separately. In the above example, 1,000 = M, 900 = CM, and 3 = III. Therefore, 1903 = MCMIII.

-- Source: Wikipedia (http://en.wikipedia.org/wiki/Roman_numerals)

**Test Input**:
------------- 
glob is I
prok is V
pish is X
tegj is L
glob glob Silver is 34 Credits  
glob prok Gold is 57800 Credits  
pish pish Iron is 3910 Credits  
how much is pish tegj glob glob ?  
how many Credits is glob prok Silver ?  
how many Credits is glob prok Gold ?  
how many Credits is glob prok Iron ?  
how much wood could a woodchuck chuck if a woodchuck could chuck wood ?

**Expecting Output**:
---------------  
pish tegj glob glob is 42  
glob prok Silver is 68 Credits  
glob prok Gold is 57800 Credits  
glob prok Iron is 782 Credits  
I have no idea what you are talking about

**Requirements**

-   Input to your program consists of lines of text detailing your notes on the conversion between intergalactic units and roman numerals. You are expected to handle invalid queries appropriately.

-   Basic components such as Input/Output processing and business logic core are essential. But please design your application as complete and modular as possible, following a standard software system requirements such as error handling, etc.

-   Well refactor the code to follow good code practices. (Some good relevant resources can be found here: http://en.wikipedia.org/wiki/Clean_Code)

-   Unit/Integration tests covering all the scenarios are good to have to show proper testing skills.

-   Implement your application using Java or Go.

-   Please document clearly about your system design solution (feel free to use professional

    diagrams if you think necessary), nontrivial assumptions as well as the clear instructions on compiling and invoking the program with program environment details. Proper server setup instructions are also essential if you develop a web application for presenting the results.

    • **NOTE**:

    -   Please DO NOT copy the solution online. We have already collected a list of typical

        solutions online. If we found your work is exactly same with some of them, we may have to

        terminate the review process.

    -   This code challenge is mean to demonstrate your software engineering skills and practice so

        Please structure your project code clean and clearly. Submission


-   You are expected to submit your source code project with the README document containing the required content stated above.

-   Create a repository on Github or Bitbucket and commit your work there. Please submit the public link of the repository containing the required source to us after you finished. We are willing to have a look on how you build the application step by step.