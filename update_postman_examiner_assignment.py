import json

with open("postman/postman_collection.json", "r") as f:
    data = json.load(f)

new_folder = {
    "name": "ExaminerAssignment Management",
    "item": [
        {
            "name": "Get List Examiner Assignments",
            "request": {
                "method": "GET",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/examiner-assignments?page=1&limit=10&exam_room_id=&examiner_id=&role=",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "examiner-assignments"],
                    "query": [
                        {"key": "page", "value": "1"},
                        {"key": "limit", "value": "10"},
                        {"key": "exam_room_id", "value": ""},
                        {"key": "examiner_id", "value": ""},
                        {"key": "role", "value": ""}
                    ]
                }
            }
        },
        {
            "name": "Create Examiner Assignment",
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
                    "raw": "{\n    \"examiner_id\": 1,\n    \"exam_room_id\": 1,\n    \"role\": \"Primary\"\n}"
                },
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/examiner-assignments",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "examiner-assignments"]
                }
            }
        },
        {
            "name": "Update Examiner Assignment",
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
                    "raw": "{\n    \"examiner_id\": 1,\n    \"exam_room_id\": 1,\n    \"role\": \"Secondary\"\n}"
                },
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/examiner-assignments/1",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "examiner-assignments", "1"]
                }
            }
        },
        {
            "name": "Delete Examiner Assignment",
            "request": {
                "method": "DELETE",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/examiner-assignments/1",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "examiner-assignments", "1"]
                }
            }
        },
        {
            "name": "Get Examiners in Room",
            "request": {
                "method": "GET",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/exam-rooms/1/examiners",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "exam-rooms", "1", "examiners"]
                }
            }
        }
    ]
}

exists = False
for item in data['item']:
    if item['name'] == 'ExaminerAssignment Management':
        item['item'] = new_folder['item']
        exists = True
        break

if not exists:
    data['item'].append(new_folder)

with open("postman/postman_collection.json", "w") as f:
    json.dump(data, f, indent=4, ensure_ascii=False)
