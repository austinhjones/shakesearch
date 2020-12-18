import React, { Component } from "react";
import axios from "axios";
import { Form, Button, Row, Container, Col, Card} from 'react-bootstrap';

let endpoint = "http://localhost:3001";

class ShakeSite extends Component {
  constructor(props) {
    super(props);

    this.state = {
      searchString: '',
      searchResults: [['shakesearch',1]],
      selectedLineContext: ''
    };

  }

    onChangeSearchString = (e) => {
        e.preventDefault();
        this.setState({searchString: e.target.value})
    }

    onSubmit = (e) => {
      e.preventDefault();
      axios.post(endpoint + "/search", null, 
      {
          params: {q: this.state.searchString}
      }).then((response) => { 
          console.log(response);
          this.setState({searchResults: response.data});
        })
    };

    displayList = () =>
        this.state.searchResults.map(result => 
                                                    <Button 
                                                        className='m-1' 
                                                        variant="light" 
                                                        size="sm" 
                                                        onClick={() => {this.searchResultOnClick(result[1])}}
                                                    >
                                                        {this.withNewlines(result[0])}
                                                    </Button>
                                                )
    
    searchResultOnClick = (lineNumber) => {
        axios.post(endpoint + "/getLineContext", null, 
        {
            params: {q: lineNumber}
        }).then((response) => { 
            console.log(response);
            this.setState({selectedLineContext: response.data});
          })
    }
            
    withNewlines = (text) => text.split('\n').map(i => {
        return <p>{i}</p>
    });

  render() {
    return (
        <Container fluid>
            <Row className="show-grid">
                <Col
                    style={{'max-height': 'calc(100vh)', 'overflow-y': 'auto'}}

                >
                    <Row className="m-1">
                        <Container fluid>
                            <Form onSubmit={this.onSubmit}>
                                <Form.Group controlId="searchbar">
                                    <Form.Label>Shakesearch!</Form.Label>
                                    <Form.Control 
                                        value={this.state.searchString}
                                        type="text" 
                                        placeholder="What art thee looking f'r" 
                                        onChange={this.onChangeSearchString}
                                    />
                                </Form.Group>
                            </Form>
                        </Container>
                    </Row>
                    {this.displayList()}

                </Col>

                <Col
                    style={{'max-height': 'calc(100vh)', 'overflow-y': 'auto'}}
                >
                    <Card fluid className="m-1">
                        <Card.Body>
                            {this.withNewlines(this.state.selectedLineContext)}
                        </Card.Body>    
                    </Card>
                </Col>
            </Row>
            
        </Container>
    );
  }
}

export default ShakeSite;