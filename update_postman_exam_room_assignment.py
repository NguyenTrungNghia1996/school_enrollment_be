import json

with open("postman/postman_collection.json", "r") as f:
    data = json.load(f)

new_folder = {
    "name": "ExamRoomAssignment Management",
    "item": [
        {
            "name": "Get List Assignments",
            "request": {
                "method": "GET",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/exam-room-assignments?page=1&limit=10&keyword=&admission_period_id=1&exam_room_id=1",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "exam-room-assignments"],
                    "query": [
                        {"key": "page", "value": "1"},
                        {"key": "limit", "value": "10"},
                        {"key": "keyword", "value": ""},
                        {"key": "admission_period_id", "value": "1"},
                        {"key": "exam_room_id", "value": "1"}
                    ]
                }
            }
        },
        {
            "name": "Create Assignment",
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
                    "raw": "{\n    \"application_id\": 1,\n    \"exam_room_id\": 1,\n    \"seat_number\": \"A01\"\n}"
                },
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/exam-room-assignments",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "exam-room-assignments"]
                }
            }
        },
        {
            "name": "Update Assignment",
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
                    "raw": "{\n    \"exam_room_id\": 1,\n    \"seat_number\": \"A02\"\n}"
                },
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/exam-room-assignments/1",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "exam-room-assignments", "1"]
                }
            }
        },
        {
            "name": "Delete Assignment",
            "request": {
                "method": "DELETE",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/exam-room-assignments/1",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "exam-room-assignments", "1"]
                }
            }
        },
        {
            "name": "Get Applications in Room",
            "request": {
                "method": "GET",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/exam-rooms/1/applications",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "exam-rooms", "1", "applications"]
                }
            }
        }
    ]
}

exists = False
for item in data['item']:
    if item['name'] == 'ExamRoomAssignment Management':
        item['item'] = new_folder['item']
        exists = True
        break

if not exists:
    data['item'].append(new_folder)

with open("postman/postman_collection.json", "w") as f:
    json.dump(data, f, indent=4, ensure_ascii=False)
