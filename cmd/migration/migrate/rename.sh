#!/bin/bash

# 方法1: 去除文件名中的时间戳
remove_timestamp() {
    # # 迭代当前目录下的所有文件 
    for file in *; do
    # 确保是文件而不是目录
    if [ -f "$file" ]; then
        # 提取文件名和扩展名
        filename=$(basename -- "$file")
        extension="${filename##*.}"
        
        # 去除时间戳部分
        new_name=$(echo "$filename" | sed 's/^[0-9]\{14\}_//')
        
        # 重命名文件
        mv "$file" "$new_name"
        
        echo "重命名文件: $file 到 $new_name"
    fi
done
}

# 方法2: 还原文件名（需要有备份或者日志）
restore_filenames() {
    # 设置时间戳格式
    timestamp=$(date +"%Y%m%d%H%M%S")

    # 创建日志文件
    log_file="rename_log.txt"
    echo "=== 重命名日志 ===" > "$log_file"

    # 迭代当前目录下以 "00" 开头的文件
    for file in 0*; do
        # 确保是文件而不是目录
        if [ -f "$file" ]; then
            # 提取文件名和扩展名
            filename=$(basename -- "$file")
            extension="${filename##*.}"
            
            # 获取数字部分和文件名部分
            number=$(echo "$filename" | cut -d'_' -f1)
            name=$(echo "$filename" | cut -d'_' -f2-)
            
            # 设置时间戳格式
            timestamp=$(expr $timestamp + 10)
            # 新文件名为时间戳 + 原数字部分 + "_" + 原文件名 + 扩展名
            new_name="${timestamp}_${name}"

            # 重命名文件
            mv "$file" "$new_name"
            
            # 记录日志
            echo "重命名文件: $file 到 $new_name" >> "$log_file"
            
            # 延时一段时间
            # sleep $time_interval
        fi
    done

    echo "重命名完成，日志保存在 $log_file"
}

# 检查参数
if [ "$#" -ne 1 ]; then
    echo "错误: 请提供一个参数，可选值为 'remove_timestamp' 或 'restore_filenames'"
    exit 1
fi

# 根据参数执行相应的方法
case "$1" in
    remove_timestamp)
        remove_timestamp
        ;;
    restore_filenames)
        restore_filenames
        ;;
    *)
        echo "错误: 无效的参数，可选值为 'remove_timestamp' 或 'restore_filenames'"
        exit 1
        ;;
esac

echo "脚本执行完成"
