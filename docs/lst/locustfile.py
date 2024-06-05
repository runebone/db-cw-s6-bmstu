from faker import Faker
from locust import HttpUser, TaskSet, task, between

fake = Faker()

words = ["the", "be", "of", "and", "a", "to", "in", "he", ...]
document_ids = ["41b9023a-77d4-4d4e-aa2c-8305b937443a", ...]

class UserBehavior(TaskSet):
    @task(2) # Вес задачи - 2
    def post_search(self):
        headers = {"Content-Type": "application/json"}
        payload = {
            "content": fake.random_element(words),
        }
        self.client.post("/search", json=payload, headers=headers)

    @task(1)
    def get_register(self):
        self.client.get("/register")
    # ...

    @task(1)
    def get_search(self):
        self.client.get("/search")

    @task(1)
    def get_document(self):
        id = fake.random_element(document_ids)
        self.client.get(f"/d/{id}")
    
class WebsiteUser(HttpUser):
    tasks = [UserBehavior]
    wait_time = between(1, 5)
