import tensorflow as tf

csv_file = "D:\\gitProject\\WoodETF\\download\\20211009\\arkk.csv"

# string_input_producer? error v2?
file_queue = tf.train.string_input_producer([csv_file])  # 设置文件名队列，这样做能够批量读取文件夹中的文件

reader = tf.TextLineReader(skip_header_lines=1)  # 使用tensorflow文本行阅读器，并且设置忽略第一行
key, value = reader.read(file_queue)
defaults = [[0.], [0.], [0.], [0.], [0.], [0.], [0.], [0.], [0.]]  # 设置列属性的数据格式
LOW, AGE, LWT, RACE, SMOKE, PTL, HT, UI, BWT = tf.decode_csv(value, defaults)
# 将读取的数据编码为我们设置的默认格式
vertor_example = tf.stack([AGE, LWT, RACE, SMOKE, PTL, HT, UI])  # 读取得到的中间7列属性为训练特征
vertor_label = tf.stack([BWT])  # 读取得到的BWT值表示训练标签

# 用于给取出的数据添加上batch_size维度，以批处理的方式读出数据。可以设置批处理数据大小，是否重复读取数据，容量大小，队列末尾大小，读取线程等属性。
example_batch, label_batch = tf.train.shuffle_batch([vertor_example, vertor_label], batch_size=10, capacity=100, min_after_dequeue=10)

# 初始化Session
with tf.Session() as sess:
    coord = tf.train.Coordinator()  # 线程管理器
    threads = tf.train.start_queue_runners(coord=coord)
    print(sess.run(tf.shape(example_batch)))  # [10  7]
    print(sess.run(tf.shape(label_batch)))  # [10  1]
    print(sess.run(example_batch)[3])  # [ 19.  91.   0.   1.   1.   0.   1.]
    coord.request_stop()
    coord.join(threads)
