#!/bin/bash

while true; do
    # 获取服务状态
    status=$(systemctl is-active sublink)
    echo "当前版本: 1.6"
    # 判断服务状态并打印
    if [ "$status" = "active" ]; then
        echo "当前运行状态: 已运行"
    else
        echo "当前运行状态: 未运行"
    fi
    echo "1. 启动服务"
    echo "2. 停止服务"
    echo "3. 卸载安装"
    echo "4. 查看服务状态"
    echo "5. 查看运行目录"
    echo "6. 修改端口"  # 新增选项
    echo "0. 退出"
    echo -n "请选择一个选项: "
    read option

    case $option in
        1)  
            systemctl start sublink
            systemctl daemon-reload
            ;;
        2)
            systemctl stop sublink
            systemctl daemon-reload
            ;;
        3)
            systemctl stop sublink
            systemctl disable sublink
            rm /etc/systemd/system/sublink.service
            systemctl daemon-reload
            rm /usr/bin/sublink
            ;;
        4)
            systemctl status sublink
            ;;
        5)
            echo "运行目录: /usr/local/bin/sublink"
            echo "需要备份的目录为db,template目录为模版文件可备份可不备份"
            ;;
        6)  # 处理修改端口的选项
            systemctl stop sublink
            echo -n "请输入新的端口号: "
            read new_port
            systemctl start sublink run --port "$new_port"
            systemctl daemon-reload
            echo "服务已使用新端口 $new_port 启动"
            ;;
        0)
            exit 0
            ;;
        *)
            echo "无效的选项"
            ;;
    esac
done
