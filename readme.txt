==================================================================================================
                            MOBILE LOANS APP
==================================================================================================


 Project Structure
 =================================

 main.go     -- this is the applications start program 
 compile     -- compiles the application for the relevant platforms
 bin/        -- contains executables for each operating system
     linux/mobile_report_generator
     mac/mobile_report_generator
     windows/mobile_report_generator.exe

packages/   -- contains mobile loans custom business logic
         common/   -- shared libraries
             utilities_test.go  -- unit tests for utilies module
             utilies.go         -- functions used for parsing and cleaning 
                                   CSV data.
             fixtures/sample_loans.csv  -- sample test data
        
        models/   -- contains models to model the data in CSV
              models_test.go   -- unit tests for Network model 
              Network.go       -- Network Struct and associated methods for handling
                                  networks parsed from the CSV file.
reports/     -- contains CSV input and output files
             Loans.csv    -- place the CSV file you wish to import here 
             Output.csv   -- generated report file 

readme.txt 

How do i run this application?
=============================

Depending on the platform you are using , simpy run the binary file found in the bin directory for your 
operating system from the projects main folder in your terminal i.e. mobile_loans/ e.g. :

For Mac: 
1. Open Terminal
2. cd mobile_loans/
3. Execute: bin/mac/mobile_report_generator 

For Linux: 
1. Open Terminal
2. cd mobile_loans/
3. Execute: bin/linux/mobile_report_generator 

For Windows**: 
1. Open CMD
2. cd mobile_loans
3. Execute: bin\linux\mobile_report_generator 

** Tested on Windows 7 AMD 64bit HP device, should work fine in other versions however I haven't had sufficient
  time or windows versions to test on the various devices and windows versions. 


(Optional) You can directly build and run the app on your OS instead of using the precompiled versions :
go run main.go       (go must be installed, instructions here : https://golang.org/doc/install)
   
   NOTE: with the above approach, the app will be a few seconds slower since go needs to compile first before
         running main.go.
  


Why GoLang?
==========

While my stronger languages are PHP, Python, JavaScript , I chose GO for the following reasons :

1. Super Fast - GO outperforms Python and PHP with ease ( benchmark tests for performance comparison can be found here
                 https://tinyurl.com/ph39q9w ) . This will therefore mean that the GO version will be more efficient
                  with bigger CSV files and put less strain on the underlying hardware. 

2. Module by Design - While GO is not exactly an OOP language , it's design patterns simulate modular OOP design with 
                      conventions for packaging code into modules and structs allowing for neat well
                      structured separation of business logic.

3. Strict Convention - GO standards are very picky about how you name variables , where you put files and how you should
                       document code. This ensures all programmers working on a project write code in a uniform manner
                       which makes maintenance so much easier and more efficient.

4. Learning Experience - Two of the core creators of GO are  Rob Pike and Ken Thompson, these are guys that have contributed
                         to projects like Unix , B (the predeccessor to C) and Google and therefore amongst the best
                        programmers in the world. GO introduces a new way of programming for the modern era and is designed
                        to build high performance applications with a cleaner syntax compared to something like C/C++. 

                        I therefore want to master this language and add it to my Python , PHP, JavaScript toolset; as I feel that it's got great potential therefore this posed a good learning exercise to improve my GO skills.  

5. Power of C with the beauty of Python - Go gives you high performance programs similar to C but also makes your development 
                                          workflow that much easier with a Python like syntax structure with static typing.


NOTES
=====

>> I omitted the "MSISDN" column because the Output.csv report is somewhat of a summary of totals.
>> To avoid another uneeded for loop - I have printed the "Total Revenue From ALL Networks" on every networks line item. This 
   will also make parsing the Output.csv a bit easier. 
>> The outputted CSV utilizes quotes '' for data fields and commas "," to seperate columns.     
>> Full documentation provided via comments throughout the source code.
>> While there may seem more folders than needed , I designed this solution around an "MVC" design pattern therefore 
   providing a "mini framework" for the application to grow and enforce an organised structure to follow when adding
   further features in future. The go compile will optimize and merge all these into a single binary. 
>> to run unit tests simply issue the following command in the projects root folder in your terminal :
    go test ./...
