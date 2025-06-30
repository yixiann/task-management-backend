# CVWO Task Management Assignment

## https://task-management-yixiann.vercel.app/overview

## Use Case:

The main purpose of this task management system is to provide a platform for corporations to efficiently manage taskings. Taskings can be created, read, updated, and deleted. To further aid in completion of tasks, a status, priority and deadline column is added in the datatable. This will enable users to sort and filter taskings based on status and priority. A tagging system will also be included to allow users to sort taskings based on category.

To improve user experience, the tags will be coloured and users can update these at any time. Tagging names changed will also update taskings. Popup boxes will also warn users if they are about to perform irreversible actions such as deleting a task. Loading icons and spinners will also be implemented to increase responsiveness of the page. Settings will also allow users to choose their desired language.

Additionally, users will be able to export taskings in a CSV file for dissemination. In future an import csv file can be implemented where taskings can be easily uploaded. A calendar view will also allow users to gain a more holistic overview of the taskings in-hand.

We hope to improve overall efficiency and productivity through the implementation of this system!

## What was used

github.com/gin-gonic/gin
github.com/go-gorp/gorp
github.com/go-sql-driver/mysql

## Front End

## https://github.com/yixiann/task-management

## Running Locally

Clone the repository

```sh
git clone https://github.com/yixiann/task-management-backend
cd task-management-backend
go build -o bin/task-management-backend -v .

export HOST=localhost
export USER=your_db_user
export PASSWORD=your_db_password
export DBNAME=your_db_name

go run main.go

```

## Created By

Tan Yi Xian
<br />A0233317M
<br />First-Year Undergraduate in NUS Double Degree Program
<br />(Computer Science and Business Administration AY2021/2022)
