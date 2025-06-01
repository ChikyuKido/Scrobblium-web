import random
import time
import requests
import json

endpoint = "http://localhost:8080/api/v1/track/add"
token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTEzNjc2ODgsInVzZXJuYW1lIjoiYWRtaW4ifQ.D8s-PtHPXkTvcScHAd3wYJcfCHH3UtZanE6m1Q77WhY"

artists_pool = ["Radiohead", "Nirvana", "Kendrick Lamar", "Daft Punk", "Björk"]
albums_pool = [
    {"name": "In Rainbows", "artists": ["Radiohead"]},
    {"name": "Nevermind", "artists": ["Nirvana"]},
    {"name": "To Pimp a Butterfly", "artists": ["Kendrick Lamar"]},
    {"name": "Discovery", "artists": ["Daft Punk"]},
    {"name": "Post", "artists": ["Björk"]}
]
tracks_pool = [
    "Weird Fishes", "Smells Like Teen Spirit", "Alright", "Harder Better Faster Stronger", "Army of Me"
]

now = int(time.time())
end = now - (6 * 30 * 24 * 60 * 60)  # ~6 months in seconds

headers = {
    'Content-Type': 'application/json',
    'Authorization': f'Bearer {token}'
}

count = 10000
failures = 0
start_time = time.time()

for i in range(count):
    album = random.choice(albums_pool)
    track_data = {
        "max_progress": random.randint(180, 300),
        "progress": random.randint(60, 300),
        "time_listened": random.randint(30, 300),
        "name": random.choice(tracks_pool),
        "artists": album["artists"],
        "listened_at": random.randint(end, now),
        "album": {
            "name": album["name"],
            "artists": album["artists"]
        }
    }

    response = requests.post(endpoint, headers=headers, data=json.dumps(track_data))
    if response.status_code != 200:
        failures += 1

end_time = time.time()
total_time = end_time - start_time
avg_time = total_time / count
rps = count / total_time

print(f"Total time: {total_time:.2f} seconds")
print(f"Average time per request: {avg_time:.4f} seconds")
print(f"Requests per second: {rps:.2f}")
print(f"Failures: {failures}")
