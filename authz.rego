package example.authz

default allow = false

allow {
    is_admin
}

is_admin {
    input.sync.groups[_] = "admin"
}