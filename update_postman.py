import json

with open("postman/postman_collection.json", "r") as f:
    data = json.load(f)

subject_folder = {
    "name": "Subject Management",
    "item": [
        {
            "name": "Get Public Subjects",
            "request": {
                "method": "GET",
                "header": [],
                "url": {
                    "raw": "{{base_url}}/api/v1/public/subjects",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "public", "subjects"]
                }
            }
        },
        {
            "name": "Get List Subjects (Admin)",
            "request": {
                "method": "GET",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/subjects?page=1&limit=10&keyword=",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "subjects"],
                    "query": [
                        {"key": "page", "value": "1"},
                        {"key": "limit", "value": "10"},
                        {"key": "keyword", "value": ""}
                    ]
                }
            }
        },
        {
            "name": "Get Subject Detail (Admin)",
            "request": {
                "method": "GET",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/subjects/1",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "subjects", "1"]
                }
            }
        },
        {
            "name": "Create Subject",
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
                    "raw": "{\n    \"code\": \"TOAN\",\n    \"name\": \"Toán Học\",\n    \"display_order\": 1,\n    \"is_active\": true\n}"
                },
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/subjects",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "subjects"]
                }
            }
        },
        {
            "name": "Update Subject",
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
                    "raw": "{\n    \"code\": \"TOAN\",\n    \"name\": \"Toán Học Nâng Cao\",\n    \"display_order\": 2,\n    \"is_active\": true\n}"
                },
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/subjects/1",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "subjects", "1"]
                }
            }
        },
        {
            "name": "Update Subject Status",
            "request": {
                "method": "PATCH",
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
                    "raw": "{\n    \"is_active\": false\n}"
                },
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/subjects/1/status",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "subjects", "1", "status"]
                }
            }
        }
    ]
}

exists = False
for item in data['item']:
    if item['name'] == 'Subject Management':
        item['item'] = subject_folder['item']
        exists = True
        break

if not exists:
    data['item'].append(subject_folder)

with open("postman/postman_collection.json", "w") as f:
    json.dump(data, f, indent=4, ensure_ascii=False)
