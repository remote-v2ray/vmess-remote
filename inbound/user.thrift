namespace go user

struct Account {
  1: required string ID;
  2: required i32 AlterID;
}

struct MemoryUser {
  1: required string Email;
  2: required i32 Level;
  3: required Account Account;
}

struct GetUserResponse {
  1: required MemoryUser User;
  2: required i64 Time;
  3: required bool Found;
}

service UserSvc {
  void Add(1: MemoryUser u);
  void Close();
  GetUserResponse Get(1: binary userHash);
  void Remove(1: string email);
}
