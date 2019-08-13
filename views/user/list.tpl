<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>User List</title>
</head>
<body>
    <h2>User List</h2>
    {{range .users}}
    <p>
        ID为：{{.Id}}，用户名为：{{.Username}}
    </p>
    {{end}}

    {{range $index,$user := .users}}
    <p>
        Index为：{{$index}}，ID为：{{$user.Id}}，用户名为：{{$user.Username}}
    </p>
    {{end}}

</body>
</html>