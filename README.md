# GO + GraphQL

#### sets list
```/set?query={list{date,data{id,date,repeats{id,weight,count},exercise{id,name,category{id,name}}}}}```

#### exercises list
```/exercise?query={list{id,name,category{id,name}}}```

#### categories list
```/category?query={list{id,name}}```

#### category by id
```/category?query={category(id:1){id,name}}```
