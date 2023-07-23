from locust import HttpUser, task

'''
Install

     sudo apt install pip
     pip3 install locust

To Run:

     locust -f ./locustfile.py --host=http://localhost:3080
     locust -f ./locustfile.py --host=http://127.0.0.1:8000

Then open browser (default link):

     http://0.0.0.0:8089

'''
class QuickstartRandoms(HttpUser):

    @task(1)
    def index_page(self):
        self.client.get("/randoms")


'''
curl -v http://localhost:3080/randoms
wrk -t2 -d10s -c 8 http://localhost:3080/randoms
'''
