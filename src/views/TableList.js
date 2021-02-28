import React , { useEffect, useState }from "react";

import { getDataUser } from "../service/index"
// react-bootstrap components
import {
  Badge,
  Button,
  Card,
  Navbar,
  Nav,
  Table,
  Container,
  Row,
  Col,
} from "react-bootstrap";

const TableList = () => {


  const [listUsers, setUser] = useState([])

  useEffect(() => {
      getListUser()
  }, [])


  const getListUser = () => {
      getDataUser().then((response) => {
          setUser(response.data.data)
console.log("ini",response.data.data)
      }).catch((error) => {
          console.log("error",error)
          
      })
  }


  return (
    <>
      <Container fluid>
        <Row>
          <Col md="12">
            <Card className="strpied-tabled-with-hover">
              <Card.Header>
                <Card.Title as="h4">Data User</Card.Title>
               
              </Card.Header>
              <Card.Body className="table-full-width table-responsive px-0">
                <Table className="table-hover table-striped">
                  <thead>
                    <tr>
                      <th className="border-0">Username</th>
                      <th className="border-0">NIK</th>
                      <th className="border-0">Tanggal Lahir</th>
                      <th className="border-0">Profesi</th>
                      <th className="border-0">Pendidikan</th>
                    </tr>
                  </thead>
                  <tbody>

                     {listUsers?.map(listUser  => (
                        <tr >
                            <td> {listUser.Username} </td>
                            <td> {listUser.IdentityNumber} </td>
                            <td> {listUser.DateOfBirth} </td>
                            <td> {listUser.Profession.ProfessionName} </td>
                            <td> {listUser.Education.EducationName} </td>
                         </tr>
                    ))
                    }
                 
                  </tbody>
                </Table>
              </Card.Body>
            </Card>
          </Col>
   
        </Row>
      </Container>
    </>
  );
}

export default TableList;
