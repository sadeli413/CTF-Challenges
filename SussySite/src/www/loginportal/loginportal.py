#!/usr/bin/env python3
import os
from flask import Flask, request

app = Flask(__name__)


class User:
    def __init__(self, name, password):
        self.name = name
        self.password = password

    def __repr__(self):
        if (self.name == "admin") and (self.password == os.environ["_SECRET"]):
            return "Welcome, admin!"
        else:
            return "Welcome " + self.name + ", you are a temporary user."


@app.route('/', methods=["GET", "POST"])
def Login():
    if request.method == "POST":
        u = User(
            request.values.get("username"),
            request.values.get("password")
        )

        # TODO: Format strings shouldn't be used in python anymore
        greeting = "Logging in as " + u.name + "...\n\n{user}"
        return greeting.format(user=u)
    else:
        greeting = ("Hello! Please Submit your username and password.\n"
                    "Example:\n"
                    "curl -X POST 'localhost:5000' -d 'username=sadeli&password=supersecret'\n")
        return greeting


if __name__ == "__main__":
    app.run()
