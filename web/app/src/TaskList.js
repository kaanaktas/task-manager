import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

let endpoint = "http://localhost:8080/api/task";

class TaskList extends Component {
    constructor(props) {
        super(props);

        this.state = {
            task: "",
            items: []
        };
    }

    componentDidMount() {
        this.getTask();
    }

    onChange = event => {
        this.setState({
            [event.target.name]: event.target.value
        });
    };

    onSubmit = () => {
        let { task } = this.state;
        if (task) {
            axios
                .post(
                    endpoint,
                        `name=${task}`,
                    {
                        headers: {
                            "Content-Type": "application/x-www-form-urlencoded",
                        },
                    }
                )
                .then(res => {
                    this.getTask();
                    this.setState({
                        task: ""
                    });
                    console.log(res);
                });
        }
    };

    getTask = () => {
        axios.get(endpoint).then(res => {
            console.log(res);
            if (res.data) {
                this.setState({
                    items: res.data.map(item => {
                        let color = "yellow";

                        if (item.done) {
                            color = "green";
                        }
                        return (
                            <Card key={item.id} color={color} fluid>
                                <Card.Content>
                                    <Card.Header textAlign="left">
                                        <div style={{ wordWrap: "break-word" }}>{item.name}</div>
                                    </Card.Header>

                                    <Card.Meta textAlign="right">
                                        <Icon
                                            name="check circle"
                                            color="green"
                                            onClick={() => this.updateTask(item.id)}
                                        />
                                        <span style={{ paddingRight: 10 }}>Done</span>
                                        <Icon
                                            name="delete"
                                            color="red"
                                            onClick={() => this.deleteTask(item.id)}
                                        />
                                        <span style={{ paddingRight: 10 }}>Delete</span>
                                    </Card.Meta>
                                </Card.Content>
                            </Card>
                        );
                    })
                });
            } else {
                this.setState({
                    items: []
                });
            }
        });
    };

    updateTask = id => {
        axios
            .put(endpoint + "/" + id)
            .then(res => {
                console.log(res);
                this.getTask();
            });
    };

    deleteTask = id => {
        axios
            .delete(endpoint + "/" + id)
            .then(res => {
                console.log(res);
                this.getTask();
            });
    };
    render() {
        return (
            <div>
                <div className="row">
                    <Header className="header" as="h2">
                        TASK LIST
                    </Header>
                </div>
                <div className="row">
                    <Form onSubmit={this.onSubmit}>
                        <Input
                            type="text"
                            name="task"
                            onChange={this.onChange}
                            value={this.state.name}
                            fluid
                            placeholder="Create Task"
                        />
                        {/* <Button >Create Task</Button> */}
                    </Form>
                </div>
                <div className="row">
                    <Card.Group>{this.state.items}</Card.Group>
                </div>
            </div>
        );
    }
}

export default TaskList;