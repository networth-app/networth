import React from "react";
import {
  Card,
  CardBody,
  CardHeader,
  CardTitle,
  Table,
  Row,
  Col,
  ListGroup,
  ListGroupItem,
} from "reactstrap";
import { get } from "../../helpers/helpers.js";

class Accounts extends React.Component {
  constructor(props) {
    super(props);
    this.state = { accounts: {} };
  }

  async componentDidMount() {
    try {
      const res = await get(`/accounts`);
      const body = await res.json();
      this.setState({ loading: false, accounts: body.data });
      console.log(body.data);
    } catch (e) {
      this.alert('Cannot get accounts. Problem connecting to REST API.');
      this.setState({ loading: false });
    }
  }

  render() {
    return (
      <div className="content">
        <Row>
          <Col xs={12}>
            <Card>
              <CardHeader>
                <CardTitle tag="h4">Accounts</CardTitle>
              </CardHeader>
              <CardBody>

              {Object.keys(this.state.accounts).map((key, _) => {
                return (
                  <div>
                <h6>{this.state.accounts[key].institution_name}</h6>
                <ListGroup>
                  {this.state.accounts[key].accounts.map((account, _) => {
                    return (<ListGroupItem>{account.name} - {account.official_name}</ListGroupItem>)
                  })}
                </ListGroup>
                <br />
                </div>
                )
              })}

              </CardBody>
            </Card>
          </Col>
        </Row>
      </div>
    );
  }
}

export default Accounts;