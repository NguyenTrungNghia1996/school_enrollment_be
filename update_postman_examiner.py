import json

with open("postman/postman_collection.json", "r") as f:
    data = json.load(f)

new_folder = {
    "name": "Examiner Management",
    "item": [
        {
            "name": "Get List Examiners",
            "request": {
                "method": "GET",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/examiners?page=1&limit=10&keyword=",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "examiners"],
                    "query": [
                        {"key": "page", "value": "1"},
                        {"key": "limit", "value": "10"},
                        {"key": "keyword", "value": ""}
                    ]
                }
            }
        },
        {
            "name": "Get Examiner Detail",
            "request": {
                "method": "GET",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/examiners/1",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "examiners", "1"]
                }
            }
        },
        {
            "name": "Create Examiner",
            "request": {
                "method": "POST",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    },
                    {
                        "key": "Content-Type",
                        "value": "application/json"
                    }
                ],
                "body": {
                    "mode": "raw",
                    "raw": "{\n    \"full_name\": \"Ngô Cán Bộ\",\n    \"organization_name\": \"Trường THPT A\",\n    \"phone_number\": \"0988777666\"\n}"
                },
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/examiners",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "examiners"]
                }
            }
        },
        {
            "name": "Update Examiner",
            "request": {
                "method": "PUT",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    },
                    {
                        "key": "Content-Type",
                        "value": "application/json"
                    }
                ],
                "body": {
                    "mode": "raw",
                    "raw": "{\n    \"full_name\": \"Ngô Cán Bộ Cập Nhật\",\n    \"organization_name\": \"Trường THPT A\",\n    \"phone_number\": \"0988777666\"\n}"
                },
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/examiners/1",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "examiners", "1"]
                }
            }
        },
        {
            "name": "Delete Examiner",
            "request": {
                "method": "DELETE",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/examiners/1",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "examiners", "1"]
                }
            }
        }
    ]
}

exists = False
for item in data['item']:
    if item['name'] == 'Examiner Management':
        item['item'] = new_folder['item']
        exists = True
        break

if not exists:
    data['item'].append(new_folder)

with open("postman/postman_collection.json", "w") as f:
    json.dump(data, f, indent=4, ensure_ascii=False)
