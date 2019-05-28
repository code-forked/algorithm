//二叉树的节点对象:用来保存数据，以及和其他节点的链接
function Node(data, left, right) {
    this.data = data;
    this.left = left;
    this.right = right;
}
Node.prototype = {
    constructor: Node,
    show(){
        return this.data;
    }
}

//二叉树:包含一个数据成员（根节点Node对象）
/**
 * 插入节点方法insert：
 * 1 创建一个Node对象，将数据传入该对象保存
 * 2 检查BST是否有根节点，如果没有，那么是棵新树，该节点即使根节点
 * 3 如果待插入节点不是根节点，那么需要遍历BST，找到插入的适当位置
 * 
 * 遍历节点方法：
 * 遍历方式有三种：中序，先序，后序，
 * 中序遍历按照节点上的健值，以升序访问BST上的所有节点，可以使用递归方式实现
 * 先序遍历先访问根节点，然后以同样的方式访问左子树和右子树
 * 后序遍历先访问叶节点，从左子树到右子树，再到根节点
 */
function BST() {
    this.root = null;
}
BST.prototype = {
    constructor: BST,
    insert(data){
        var n = new Node(data, null, null);
        if(this.root == null) {
            this.root = n;
        } else {
            var current = this.root;
            var parent;
            while(true) {
                parent = current;
                if(data < current.data) {
                    current = current.left;
                    if(current == null){
                        parent.left = n;
                        break;
                    }
                } else {
                    current = current.right;
                    if(current == null) {
                        parent.right = n;
                        break;
                    }
                }
            }
        }
    },
    inOrder(node){      //中序遍历
        if(node != null) {
            this.inOrder(node.left);
            console.log("inOrder:" + node.show());
            this.inOrder(node.right);
        }

    },
    preOrder(node){     //先序遍历
        if(node != null) {
            console.log("preOrder:" + node.show());
            this.preOrder(node.left);
            this.preOrder(node.right);
        }
    },
    postOrder(node) {   //后序遍历
        if(node != null){
            this.postOrder(node.left);
            this.postOrder(node.right);
            console.log("postOrder:" + node.show());
        }

    },
    getMin() {          //查找最小值 参数为树的方向，值为 null，left，right
        var current = this.root;
        while(current.left != null) {
            current = current.left;
        }
        return current.data;
    },
    getMax() {          //查找最大值
        var current = this.root;
        while(current.right != null) {
            current = current.right;
        }
        return current.data;
    },
    find(data) {        //查找给定值
        var current = this.root;
        while( current != null) {
            if(current.data == data) {
                return current;
            } else if(data < current.data) {
                current = current.left;
            } else {
                current = current.right;
            }
        }
        return null;
    },
    remove(data) {     
        var nowRoot = this.root;                                         //删除节点方法
        nowRoot = this.removeNode(nowRoot, data);
    },
    removeNode(node,data){
        if(node == null) {
            return null;
        }
        if(data == node.data) {

            if(node.left == null && node.right == null) {       //没有子节点的节点
                return null;
            }

            if(node.left == null) {                             //没有左子节点
                return node.right;
            }

            if(node.right == null) {                            //没有右子节点
                return node.left;
            }

            //左右子节点都有                        
            var minNode = node.right;               //获取右树最小值或者左树最大值       
            while(minNode.left != null) {
                minNode = minNode.left
            }
            node.data = minNode.data;
            node.right = this.removeNode(node.right, minNode.data);
            return node;

        } else if (data < node.data) {

            node.left = this.removeNode(node.left, data);
            return node;

        } else {

            node.right = this.removeNode(node.right, data);
            return node;

        }
    }
}