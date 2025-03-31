# probable-system

A Go/Http server.

JWT authenticated endpoints provide access to AWS Dynamodb and S3 backend storage supporting user creation, messaging, and file upload/download.

Additional endpoints to access GTFS-RT data feeds for alerts, trip updates, and vehicle position data.

Data sets for routes, route shapes, stops, stop times, and trips are imported via CSV.

Data is imported from the files through processing functions that output Go files of public slices of data as struct literals for each data type.

# notes

go get github.com/aws/aws-sdk-go-v2/service/dynamodb
go get github.com/aws/aws-sdk-go-v2/service/s3
go get github.com/aws/aws-sdk-go-v2/service/rekognition

Type	Go Struct	Description
String	types.AttributeValueMemberS	Stores a string value.
Number	types.AttributeValueMemberN	Stores a number as a string (DynamoDB stores numbers in string format).
Boolean	types.AttributeValueMemberBOOL	Stores a boolean (true or false).
Binary	types.AttributeValueMemberB	Stores binary data ([]byte).
String Set	types.AttributeValueMemberSS	Stores a set of unique strings ([]string).
Number Set	types.AttributeValueMemberNS	Stores a set of unique numbers as strings ([]string).
Binary Set	types.AttributeValueMemberBS	Stores a set of unique binary values ([][]byte).
Map	types.AttributeValueMemberM	Stores a nested map (JSON-like object).
List	types.AttributeValueMemberL	Stores a list of attribute values (similar to an array).
Null	types.AttributeValueMemberNULL	Represents a null value.