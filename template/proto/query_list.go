package proto

var QueryList = `// auto-generated by terser-cli {{.GenerateTime}}
// proto QueryList
syntax = "proto3";
package proto;

message PageInfo {
    int32 Total = 1; // 总条数
    int32 PageIndex = 2; // 当前页
    int32 PageSize = 3; // 页码条数
}

message QueryParams {
    int32 PageIndex = 1;
    int32 PageSize = 2;
    repeated QueryField QueryList = 3;
    repeated OrderField OrderList = 4;
}

message QueryField {
    string Name = 1;
    string Value = 2;
    QueryExpress Express = 3;
}
message OrderField {
    string Name = 1;
    OrderType Value = 2;
}
enum OrderType {
    DESC = 0;
    ASC = 1;
}
enum QueryExpress {
    EQ = 0;
    LT = 1;
    GT = 2;
    LIKE = 3;
    GTE = 4;
    LTE = 5;
    NE = 6;
    IN = 7;
    IS_NULL = 8;
    FIND_IN_SET = 9;
}
`