'use strict';

import React from 'react';

export default class Writer extends React.Component {
    render() {
        return (
            <div className="panel panel-default">
                <div className="panel-heading">
                    <div className="input-group">
                        <span className="input-group-addon">タイトル</span>
                        <input type="text" className="form-control" name="title" />
                    </div>
                </div>

                <div className="panel-body">
                    <div className="row">
                        <div className="col-xs-10">
                            <h4>本文</h4>
                            <textarea name="markdown" className="form-control" name="body"></textarea>
                        </div>
                    </div>

                    <div className="row">
                        <div className="col-xs-10">
                            <div className="checkbox">
                                <label>
                                    <input type="checkbox" id="is_show" name="is_show"></input> 投稿と同時に公開する
                                </label>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="panel-footer">
                    <button type="button" className="btn btn-primary">登録</button>
                </div>
            </div>
        );
    }
}
