import React, { useContext, useState, useEffect, useRef } from "react";
import { Table, Input, Button, Popconfirm, Form } from "antd";

const EditableContext = React.createContext(null);

function EditableRow(props) {
  //编辑表格行
  // {
  //   children:[
  //     {$$typeof: Symbol(react.element), type: {…}, key: 'name', ref: null, props: {…}, …},
  //     {$$typeof: Symbol(react.element), type: {…}, key: 'age', ref: null, props: {…}, …},
  //     {$$typeof: Symbol(react.element), type: {…}, key: 'address', ref: null, props: {…}, …},
  //     {$$typeof: Symbol(react.element), type: {…}, key: 'operation', ref: null, props: {…}, …}
  //   ],
  //   className: "ant-table-row ant-table-row-level-0",
  //   data-row-key: "0"
  // }
  // console.log("EditableRowProps", props);
  let [form] = Form.useForm(); //定义表单对象
  return (
    <Form form={form} component={false}>
      <EditableContext.Provider value={form}>
        <tr {...props} />
      </EditableContext.Provider>
    </Form>
  );
}

function EditableCell({
  title,
  editable,
  children,
  dataIndex,
  record,
  handleSave,
  ...restProps
}) {
  //编辑表格单元格
  // {
  //   children:[undefined, 'Easdf'],
  //   className: "ant-table-cell",
  //   colSpan: null,
  //   dataIndex: "name",
  //   editable: true,
  //   handleSave: row => {…},
  //   record: {key: '1', name: 'Easdf', age: '32', address: 'London, Park Lane no. 1'},
  //   rowSpan: null,
  //   style: {},
  //   title: "姓名",
  // }
  // console.log("EditableCellProps", props);

  let inputRef = useRef(null);
  const [editing, setEditing] = useState(false); //定义编辑状态
  const form = useContext(EditableContext);

  useEffect(() => {
    //监听编辑状态值的变化
    if (editing) {
      //如果开启编辑状态
      inputRef.current.focus(); //input输入框聚焦
    }
  }, [editing]);

  function toggleEdit() {
    //切换编辑状态
    setEditing(!editing);
    form.setFieldsValue({
      //将表格中的值赋值到表单中
      // name:Easdf
      [dataIndex]: record[dataIndex],
    });
  }

  async function save() {
    //保存事件
    try {
      const values = await form.validateFields(); //获取表单中的数据
      // {name: 'Edsdf s'}
      // console.log(values);
      toggleEdit(); //切换编辑状态
      handleSave({ ...record, ...values }); //调用保存方法
    } catch (errInfo) {
      console.log("保存失败:", errInfo);
    }
  }

  let childNode = children;
  if (editable) {
    //如果开启了表格编辑属性
    // 是否开启了编辑状态 (开启:显示输入框 关闭:显示div)
    childNode = editing ? (
      <Form.Item
        name={dataIndex}
        rules={[
          {
            required: true,
            message: `${title}是必填的.`,
          },
        ]}
      >
        <Input ref={inputRef} onPressEnter={save} onBlur={save} />
      </Form.Item>
    ) : (
      <div onClick={toggleEdit}>{children}</div>
    );
  }
  // {
  //   // {
  //   //   className: "ant-table-cell",
  //   //   colSpan: null,
  //   //   rowSpan: null,
  //   //   style: {},
  //   // }
  //   console.log(restProps)
  // }
  return <td {...restProps}>{childNode}</td>;
}
export default function AAA() {
  let tableColumns = [
    //定义表格头
    {
      title: "姓名",
      dataIndex: "name",
      with: "30%",
      editable: true,
    },
    {
      title: "年龄",
      dataIndex: "age",
    },
    {
      title: "地址",
      dataIndex: "address",
    },
    {
      title: "操作",
      dataIndex: "operation",
      render: (text, row) => {
        return (
          <Popconfirm
            title="确定要删除么？"
            onConfirm={() => {
              let arr = tableData.filter((item) => {
                return item.key !== row.key;
              });
              setTableData([...arr]);
            }}
          >
            <a>删除</a>
          </Popconfirm>
        );
      },
    },
  ];
  let [count, setCount] = useState(2);
  let [tableData, setTableData] = useState([
    //定义表格数据
    {
      key: "0",
      name: "Edward King 0",
      age: "32",
      address: "London, Park Lane no. 0",
    },
    {
      key: "1",
      name: "Edward King 1",
      age: "32",
      address: "London, Park Lane no. 1",
    },
  ]);

  tableColumns = tableColumns.map((item) => {
    //遍历表格头数组
    if (item.editable) {
      //如果表格头列开启了编辑属性
      // {title: '姓名', dataIndex: 'name', with: '30%', editable: true}
      // console.log(item);
      return {
        ...item,
        onCell: (record) => {
          // {key: '0', name: 'Edward King 0', age: '32', address: 'London, Park Lane no. 0'}
          // {key: '1', name: 'Edward King 1', age: '32', address: 'London, Park Lane no. 1'}
          // console.log(record);
          return {
            record: record,
            editable: item.editable,
            dataIndex: item.dataIndex,
            title: item.title,
            handleSave: (row) => {
              //这个方法可以获取到行编辑之后的数据
              // {key: '1', name: 'Easdf', age: '32', address: 'London, Park Lane no. 1'}
              // console.log("handleSaveRow", row);
              let findEditIndex = tableData.findIndex((item) => {
                //找到编辑行的索引
                return item.key === row.key;
              });
              let findEditObj = tableData.find((item) => {
                //找到编辑行的数据对象
                return item.key === row.key;
              });
              tableData.splice(findEditIndex, 1, { ...findEditObj, ...row }); //将最新的数据更新到表格数据中
              setTableData([...tableData]); //设置表格数据
            },
          };
        },
      };
    } else {
      // {title: '年龄', dataIndex: 'age'}
      // {title: '地址', dataIndex: 'address'}
      // {title: '操作', dataIndex: 'operation', render: ƒ}
      // console.log(item);
      return item;
    }
  });

  return (
    <div>
      <Button
        onClick={() => {
          let newData = {
            key: String(count),
            name: `Edward King ${count}`,
            age: "32",
            address: `London, Park Lane no. ${count}`,
          };
          setCount(count + 1);
          setTableData([...tableData, newData]);
        }}
        type="primary"
      >
        添加一行
      </Button>
      <Table
        bordered
        components={{
          body: {
            row: EditableRow,
            cell: EditableCell,
          },
        }}
        columns={tableColumns}
        dataSource={tableData}
      />
    </div>
  );
}
