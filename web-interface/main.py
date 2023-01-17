from flask import Flask, session, url_for, redirect, render_template, request, abort, flash, send_from_directory, send_file
import requests

openkothAPI = "localhost:8080"
app = Flask(__name__)


@app.route("/")
def root():
    return render_template("index.html")


@app.route("/rooms")
def rooms():
    r = requests.get("http://" + openkothAPI + "/rooms")
    print(r.json())
    return r.json()


@app.route("/register", methods=["GET", "POST"])
def register():
    return "Current under developement. Please come back later!"


@app.route("/login", methods=["GET", "POST"])
def login():
    return render_template("login.html")


@app.route("/info")
def info():
    return render_template("info.html")


@app.route("/about")
def about():
    return render_template("about.html")



@app.route("/tmp")
def tmp():
    return render_template("rooms.html")




if __name__ == "__main__":
    app.run(debug=True, host="0.0.0.0", port=9999)
