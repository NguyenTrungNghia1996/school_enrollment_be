import json

with open("postman/postman_collection.json", "r") as f:
    data = json.load(f)

new_folder = {
    "name": "ApplicationExamScore Management",
    "item": [
        {
            "name": "Get Application Exam Scores",
            "request": {
                "method": "GET",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/applications/1/scores",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "applications", "1", "scores"]
                }
            }
        },
        {
            "name": "Update Application Exam Scores",
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
                    "raw": "{\n    \"scores\": [\n        {\n            \"subject_id\": 1,\n            \"raw_score\": 8.5,\n            \"bonus_score\": 0.5,\n            \"final_score\": 9.0,\n            \"is_absent\": false,\n            \"notes\": \"Điểm tốt\"\n        },\n        {\n            \"subject_id\": 2,\n            \"raw_score\": 0,\n            \"bonus_score\": 0,\n            \"final_score\": 0,\n            \"is_absent\": true,\n            \"notes\": \"Vắng thi\"\n        }\n    ]\n}"
                },
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/applications/1/scores",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "applications", "1", "scores"]
                }
            }
        }
    ]
}

exists = False
for item in data['item']:
    if item['name'] == 'ApplicationExamScore Management':
        item['item'] = new_folder['item']
        exists = True
        break

if not exists:
    data['item'].append(new_folder)

with open("postman/postman_collection.json", "w") as f:
    json.dump(data, f, indent=4, ensure_ascii=False)
