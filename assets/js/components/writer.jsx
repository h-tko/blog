'use strict';

import React from 'react';
import $ from 'jquery';
import DatePicker from 'react-datepicker';
import moment from 'moment';

import 'react-datepicker/dist/react-datepicker.css';

export default class Writer extends React.Component {
    constructor(props) {
        super(props);

        this.state = this.defaultState();
    }

    defaultState() {
        return {
            title: "",
            body: "",
            release_date: null,
            is_show: 1,
        };
    }

    regist(e) {
        if (confirm("本当に登録してもよろしいですか？")) {

            var params = this.state;
            params.release_date = params.release_date.format('YYYY/MM/DD');
            console.log(params);

            $.post("/regist_blog", params, (result) => {
                if (result === "success") {
                    alert("登録しました");

                    this.setState(this.defaultState());
                }
            });
        }
    }

    render() {
        return (
            <div className="panel panel-default">
                <div className="panel-heading">
                    <div className="input-group">
                        <span className="input-group-addon">タイトル</span>
                        <input type="text" className="form-control" name="title" value={this.state.title} onChange={(e) => this.setState({title: e.target.value})} />
                    </div>
                </div>

                <div className="panel-body">
                    <div className="row">
                        <div className="col-xs-11">
                            <h4>本文</h4>
                            <textarea name="markdown" className="form-control" name="body" value={this.state.body} onChange={(e) => this.setState({body: e.target.value})}></textarea>
                        </div>
                    </div>

                    <div className="row">
                        <div className="col-xs-11">
                            <h4>公開日</h4>
                            <DatePicker selected={this.state.release_date} onChange={(date) => {this.setState({release_date: date})}} />
                        </div>
                    </div>
                </div>
                <div className="panel-footer">
                    <button type="button" className="btn btn-primary" onClick={this.regist.bind(this)}>登録</button>
                </div>
            </div>
        );
    }
}
