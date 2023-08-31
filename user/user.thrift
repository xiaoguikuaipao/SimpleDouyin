namespace go api

struct Request {
    1: string message
}

struct Response {
    1: string message
}

struct RegisterRequest {
    1: string username
    2: string password
}

struct RegisterResponse {
    1: bool OK
    2: string error
}

struct LoginRequest {
    1: string username
    2: string password
}

struct LoginResponse{
    1: bool OK
    2: string error
}

service Register {
    Response echo(1: Request req)
    RegisterResponse register(1: RegisterRequest Rreq)
}

service Login {
    Response echo(1: Request req)
    LoginResponse login(1: LoginRequest Lreq)
}