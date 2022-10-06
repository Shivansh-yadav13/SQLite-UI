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
docker pull ghcr.io/shivansh-yadav13/sqlite-ui:latest
```
**Step 2:**

Start Docker container using:
```
docker run -p { port }:{ port } -v { location of db file }:/app/sqlite_database -e SQLITE_NAME={ db file name } shivansh-yadav13/sqlite-ui:latest
```
**Optional:**
You can also use the `--name` and `-d` for the name of the container and for running in detached mode.

## **💻 Set up Project Locally**

**1.)** Fork the repository by click the `Fork` button at the top right of the page.

![Screenshot from 2022-09-22 12-07-47](https://user-images.githubusercontent.com/87603425/191675682-07be1e87-060e-4dfa-92d2-f21cb03b0de6.png)
</br>
</br>
**2.)** Navigate to a directory, open the terminal and paste the following command and replace `your-github-username` with your GitHub username:
```bash
git clone https://github.com/your-github-username/SQLite-UI.git
```
**3.)** To start the project locally we need to create a sqlite file with name as `data.db`. Navigate to the root of the project and create a folder with the name `sqlite_database` and inside it create a file `data.db`.

**4.)** Now we are ready to start the server, run the following:
```bash
go run main.go
```

**5.)** Now you can visit `localhost:3000` and you will be able to see:
![Screenshot from 2022-09-22 12-21-21](https://user-images.githubusercontent.com/87603425/191677999-c3a95f63-8f4b-4b53-8a04-f50fc4323d9d.png)

## **👥 Contributors**

[![Contributors](https://contrib.rocks/image?repo=shivansh-yadav13/sqlite-ui)](https://github.com/shivansh-yadav13/sqlite-ui/graphs/contributors)
