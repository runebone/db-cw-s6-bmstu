from faker import Faker
from locust import HttpUser, TaskSet, task, between

fake = Faker()

words = [
    "the", "be", "of", "and", "a", "to", "in", "he", "have", "it", "that",
    "for", "they", "I", "with", "as", "not", "on", "she", "at", "by", "this",
    "we", "you", "do", "but", "from", "or", "which", "one", "would", "all",
    "will", "there", "say", "who", "make", "when", "can", "more", "if", "no",
    "man", "out", "other", "so", "what", "time", "up", "go", "about", "than",
    "into", "could", "state", "only", "new", "year", "some", "take", "come",
    "these", "know", "see", "use", "get", "like", "then", "first", "any",
    "work", "now", "may", "such", "give", "over", "think", "most", "even",
    "find", "day", "also", "after", "way", "many", "must", "look", "before",
    "great", "back", "through", "long", "where", "much", "should", "well",
    "people", "down", "own", "just", "because", "good", "each", "those",
    "feel", "seem", "how", "high", "too", "place", "little", "world", "very",
    "still", "nation", "hand", "old", "life", "tell", "write", "become",
    "here", "show", "house", "both", "between", "need", "mean", "call",
    "develop", "under", "last", "right", "move", "thing", "general", "school",
    "never", "same", "another", "begin", "while", "number", "part", "turn",
    "real", "leave", "might", "want", "point", "form", "off", "child", "few",
    "small", "since", "against", "ask", "late", "home", "interest", "large",
    "person", "end", "open", "public", "follow", "during", "present",
    "without", "again", "hold", "govern", "around", "possible", "head",
    "consider", "word", "program", "problem", "however", "lead", "system",
    "set", "order", "eye", "plan", "run", "keep", "face", "fact", "group",
    "play", "stand", "increase", "early", "course", "change", "help", "line"
    ]

with open("document.ids", 'r') as file:
    document_ids = file.read().splitlines()

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

    @task(1)
    def get_login(self):
        self.client.get("/login")

    @task(1)
    def get_profile(self):
        self.client.get("/profile")

    @task(1)
    def get_search(self):
        self.client.get("/search")

    @task(1)
    def get_document(self):
        id = fake.random_element(document_ids)
        self.client.get(f"/d/{id}")
    
class WebsiteUser(HttpUser):
    tasks = [UserBehavior]
    wait_time = between(1, 5)  # Время ожидания между задачами от 1 до 5 секунд

