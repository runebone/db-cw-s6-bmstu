from faker import Faker
from locust import HttpUser, TaskSet, task, between

fake = Faker(["ru_RU"])

class UserBehavior(TaskSet):
    @task(2) # Вес задачи - 2
    def post_search(self):
        headers = {"Content-Type": "application/json"}
        payload = {
            "content": "a",
        }
        self.client.post("/search", json=payload, headers=headers)

    @task(1)
    def get_profile(self):
        self.client.get("/profile")

    @task(1)
    def get_search(self):
        self.client.get("/search")

    @task(1)
    def get_document(self):
        document_id = "41b9023a-77d4-4d4e-aa2c-8305b937443a"
        self.client.get(f"/d/{document_id}")
    
class WebsiteUser(HttpUser):
    tasks = [UserBehavior]
    wait_time = between(1, 5)  # Время ожидания между задачами от 1 до 5 секунд

