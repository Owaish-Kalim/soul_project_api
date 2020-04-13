CREATE TABLE slh_team(
    "TeamId" serial PRIMARY KEY,
    "FirstName" VARCHAR(50) NOT null,
    "LastName" VARCHAR(50) NOT null,
    "Email" VARCHAR(50) UNIQUE NOT null,
    "Password" VARCHAR(50) NOT null,
    "Address" VARCHAR(200) NOT null,
    "Token" VARCHAR(200) NOT null,
    "MobileNo" VARCHAR(10) NOT null,
    "Status" VARCHAR(50) NOT null,
    "Joining_Date" TIMESTAMP NOT null,
    "CreatedAt" TIMESTAMP NOT null
);




CREATE TABLE slh_team_has_role(
    "Team_Has_Role_Id" serial PRIMARY KEY,
    "TeamId" INT NOT null,
    "Status" VARCHAR(50) NOT null,
    "UpdatedAt" TIMESTAMP NOT null,
    "CreatedAt" TIMESTAMP NOT null
);


CREATE TABLE slh_customer(
    "Customer_Id " serial PRIMARY KEY,
    "Customer_Souls_Id" VARCHAR(50) UNIQUE NOT null,
    "Customer_Name" VARCHAR(50) NOT null,
    "Customer_Mobile_No" INT NOT null,
    "Customer_Gender" VARCHAR(50) NOT null,
    "Pincode" INT NOT null,
    "Customer_Email" VARCHAR(50) UNIQUE NOT null,
    "Customer_Address" VARCHAR(200) NOT null,
    "Status " Boolean NOT null,
    "Last_Access_Time" TIMESTAMP NOT null,
    "Registered_Source" VARCHAR(50) NOT null,
    "CreatedAt" TIMESTAMP NOT null
);



CREATE TABLE slh_customers_pending_orders(
    "Customer_Order_Id" serial PRIMARY KEY,
    "Customer_Id" INT NOT null,
    "Customer_Souls_Id" VARCHAR(50) NOT null,
    "Number_Of_Therapists_Required" INT NOT null,
    "Therapist_Gender" VARCHAR(50) NOT null,
    "Massage_For" VARCHAR(50) NOT null,
    "Slot_Time" TIMESTAMP NOT null,
    "Slot_Date" TIMESTAMP NOT null,
    "Customer_Address" VARCHAR(200) NOT null,
    "Pincode" INT NOT null,
    "Latitude" VARCHAR(200) NOT null,
    "Longitude" VARCHAR(200) NOT null,
    "Is_Order_Confirmed" Boolean NOT null,
    "Merchant_Transaction_Id" VARCHAR(50) NOT null,
    "Massage_Duration" VARCHAR(50) NOT null,
    "CreatedAt" TIMESTAMP NOT null,
    "Customer_Name" VARCHAR(50) NOT null,
    "Total_Order_Amount" INT NOT null
);



CREATE TABLE slh_transactions(
    "Customer_Order_Id" INT NOT null,
    "Customer_Id" INT NOT null,
    "Customer_Souls_Id" VARCHAR(50) NOT null,
    "Number_Of_Therapist_Required" INT NOT null,
    "Therapist_Gender" VARCHAR(50) NOT null,
    "Massage_For" VARCHAR(50) NOT null,
    "Slot_Time" TIMESTAMP NOT null,
    "Slot_Date" TIMESTAMP NOT null,
    "Customer_Address" VARCHAR(200) NOT null,
    "Pincode" INT NOT null,
    "Latitude" VARCHAR(50) NOT null,
    "Longitude" VARCHAR(50) NOT null,
    "Merchant_Transaction_Id" VARCHAR(50) NOT null,
    "Massage_Duration" VARCHAR(50) NOT null,
    "CreatedAt" TIMESTAMP NOT null,
    "Customer_Name" VARCHAR(50) NOT null,
    "Total_Order_Amount" INT NOT null,
    "Payment_Gateway_Id" VARCHAR(50) NOT null,
    "Payment_Gateway_Mode" VARCHAR(50) NOT null,
    "Transaction_Mode" VARCHAR(50) NOT null,
    "Bank_Type" VARCHAR(50) NOT null
);




CREATE TABLE slh_partners(
    "Partner_Id" serial PRIMARY KEY,
    "Partner_Name" VARCHAR(100) NOT null,
    "Partner_Gender" VARCHAR(50) NOT null,
    "Partner_Address" VARCHAR(200) NOT null,
    "Partner_Mobile_No" INT NOT null,
    "Partner_Email" VARCHAR(50) UNIQUE NOT null,
    "Pincode" INT NOT null,
    "Latitude" VARCHAR(50) NOT null,
    "Longitude" VARCHAR(50) NOT null,
    "Onboard_Date" TIMESTAMP NOT null,
    "CreatedAt" TIMESTAMP NOT null,
    "UpdatedAt" TIMESTAMP NOT null,
    "Last_Updated_By" VARCHAR(50) NOT null,
    "Per_Visit_Price_Commission" INT NOT null,
    "Commission_Type" VARCHAR(50) NOT null,
    "CreatedBy" VARCHAR(50) NOT null
);


CREATE TABLE slh_assign_customer_with_partner(
    "Customer_Souls_Id" VARCHAR(50) NOT null,
    "Customer_Name" VARCHAR(50) NOT null,
    "Customer_Id" INT NOT null,
    "Merchant_Transaction_Id" VARCHAR(50) NOT null,
    "Status" VARCHAR(50) NOT null,
    "Commission_Amount" INT NOT null,
    "Created_By" VARCHAR(50) NOT null,
    "Updated_By" VARCHAR(50) NOT null,
    "Slot_Date" TIMESTAMP NOT null,
    "Slot_Time" TIMESTAMP NOT null,
    "CreatedAt" TIMESTAMP NOT null
);