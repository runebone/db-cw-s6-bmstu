from locust import HttpUser, TaskSet, task, between

class UserBehavior(TaskSet):
    @task(2) # Вес задачи - 2
    def post_data(self):
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
        document_id = "52a7a3e2-e086-47d7-80ef-3f10524d5b3a"
        self.client.get(f"/d/{document_id}")
    
class WebsiteUser(HttpUser):
    tasks = [UserBehavior]
    wait_time = between(1, 5)  # Время ожидания между задачами от 1 до 5 секунд

