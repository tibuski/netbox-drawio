curl -X POST ^
-H "Content-Type: application/json" ^
-H "Accept: application/json; indent=4" ^
https://demo.netbox.dev/api/users/tokens/provision/ ^
--data "{\"username\": \"admin\",\"password\": \"admin\"}"