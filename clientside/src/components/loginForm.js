import React, { Component } from 'react';

import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField'
import { red } from '@material-ui/core/colors';

const styles = {
    root: {
        flexGrow: 1,
    },
    grow: {
        flexGrow: 1,
    },
    block: {
        display: "block",
    },
    errormessage: {
        color: red,
    },
};

class LoginFormComponent extends Component {
    constructor(props) {
        super(props);
        this.state = {
            username: '',
            password: '',
            error: false
        };
    };
    handleChange = (input) => event => {
        this.setState({ [input]: event.target.value, error: null })
    };
    handleSubmit = (event) => {
        fetch('/api/auth/login', { method: 'POST', body: JSON.stringify({ username: this.state.username, password: this.state.password }) })
            .then(data => data.json())
            .then(res => {
                if (!res.Success) {
                    //set error status to true, makes form red
                    this.setState({error: true})
                }
                else {
                    //store token and send on to dashboard
                    localStorage.setItem('token', res.Data.token)
                    const { history } = this.props
                    history.push('/dashboard')
                }
            }
            );

        event.preventDefault();
    };

    render() {
        const { classes } = this.props;

        return (
            <form onSubmit={this.handleSubmit}>

                <h2>Login</h2>
                <TextField
                    error={this.state.error}
                    id="login-username"
                    label="Username"
                    placeholder="Username"
                    className={classes.textField}
                    margin="normal"
                    variant="outlined"
                    onChange={this.handleChange('username')}
                />
                <br></br>

                <TextField
                    error={this.state.error}
                    id="login-password"
                    label="Password"
                    className={classes.textField}
                    type="password"
                    autoComplete="current-password"
                    margin="normal"
                    variant="outlined"
                    onChange={this.handleChange('password')}
                />
                <br></br>

                <Button variant="outlined" className={classes.button} label="submit" type="submit">
                    Login
                </Button>

                <p hidden={!this.state.error} className={classes.errormessage}>unable to authenticate</p>
            </form>
        )
    }
}

LoginFormComponent.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(LoginFormComponent);