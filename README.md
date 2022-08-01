# 🌐 SQLite UI
**SQLite UI** is a web user interface for SQLite database built with Golang.
It is built in Golang using the **Fiber** package.
**Fiber** is an [Express](https://github.com/expressjs/express) inspired **web framework**
You can read more about it **[here](https://github.com/gofiber/fiber)**.
Project is right now at it's very intial stage and under developement. Our plan is to make this a **Good Open Source** project, where people can easily start contributing, fix bugs and add more features 🙂.

## **🔧 Usage**

**Step 1:**

Pull Docker Image.
```bash
docker pull shivanshyadav/sqlite_ui
```
**Step 2:**

Start Docker container using:
```
docker run -p { port }:{ port } -v { location of db file }:/app/sqlite_database -e SQLITE_NAME={ db file name } shivanshyadav/sqlite_ui
```
**Optional:**
You can also use the `--name` and `-d` for the name of the container and for running in detached mode.
