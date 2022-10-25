![SQLite-UI](https://socialify.git.ci/shivansh-yadav13/sqlite-ui/image?description=1&font=Raleway&forks=1&issues=1&logo=https%3A%2F%2Fdwglogo.com%2Fwp-content%2Fuploads%2F2018%2F03%2FSQLite_Vector_logo-1024x705.png&pattern=Plus&pulls=1&stargazers=1&theme=Light)

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
docker run -p 3000:3000 -v { location of db file }:/app -e SQLITE_NAME=/app/{ db file name } shivansh-yadav13/sqlite-ui:latest
```
**Optional:**
You can also use the `--name` and `-d` for the name of the container and for running in detached mode.

## **💻 Set up Project Locally**

**1.)** Fork the repository by click the `Fork` button at the top right of the page.

![Screenshot from 2022-09-22 12-07-47](https://user-images.githubusercontent.com/87603425/191675682-07be1e87-060e-4dfa-92d2-f21cb03b0de6.png)

**2.)** Navigate to a directory, open the terminal and paste the following command and replace `your-github-username` with your GitHub username:
```bash
git clone https://github.com/your-github-username/SQLite-UI.git
```

**3.)** Build the TailwindCSS bundle, see https://tailwindcss.com/docs/installation for instruction to install the cli.

```bash
./tailwindcss -i ./static/css/app.css -o ./static/css/build.min.css --minify
```

**4.)** Now we are ready to start the server, run the following:

```bash
$ go run main.go
```

This will create an empty `data.db` in the current directory. Alternatively, you can also pass in an existing database file as follow:

```bash
$ SQLITE_NAME=<path-to-db-file> go run main.go
```

**5.)** Now you can visit `localhost:3000` and you will be able to see:

![Screenshot from 2022-09-22 12-21-21](https://user-images.githubusercontent.com/87603425/191677999-c3a95f63-8f4b-4b53-8a04-f50fc4323d9d.png)

## **👥 Contributors**

[![Contributors](https://contrib.rocks/image?repo=shivansh-yadav13/sqlite-ui)](https://github.com/shivansh-yadav13/sqlite-ui/graphs/contributors)
