import json

with open("postman/postman_collection.json", "r") as f:
    data = json.load(f)

new_folder = {
    "name": "Dashboard",
    "item": [
        {
            "name": "Get Dashboard Summary",
            "request": {
                "method": "GET",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/dashboard/summary?admission_period_id=",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "dashboard", "summary"],
                    "query": [
                        {"key": "admission_period_id", "value": ""}
                    ]
                }
            }
        }
    ]
}

exists = False
for item in data['item']:
    if item['name'] == 'Dashboard':
        item['item'] = new_folder['item']
        exists = True
        break

if not exists:
    data['item'].append(new_folder)

with open("postman/postman_collection.json", "w") as f:
    json.dump(data, f, indent=4, ensure_ascii=False)
