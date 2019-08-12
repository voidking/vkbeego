<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Test</title>
</head>
<body>
    <h2>Test Page</h2>
    <p>Website: {{.Website}}</p>
    <p>Email: {{.Email}}</p>
    <p>
        Gender:
        {{if .IsMale}}
        Male
        {{else}}
        Female
        {{end}}
    </p>
    <p>
        Lucky Numbers:
        {{range .LuckyNumbers}}
        {{.}}&nbsp;&nbsp;
        {{end}}
    </p>
</body>
</html>