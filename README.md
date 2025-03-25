# probable-system

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