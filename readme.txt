==================================================================================================
                            MOBILE LOANS APP
==================================================================================================

A simple Go based app to parse CSV files and extract network transactional information.

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
