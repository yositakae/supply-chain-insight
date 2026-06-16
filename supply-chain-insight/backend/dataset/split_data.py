import pandas as pd
import random
print("กำลังอ่านไฟล์ DataCo... (อาจใช้เวลาสักครู่)")
df = pd.read_csv('DataCoSupplyChainDataset.csv', encoding='latin1')

# ---------------------------------------------------------
# 1. สร้าง users.csv (เหมือนเดิม)
# ---------------------------------------------------------
print("กำลังสร้าง users.csv...")
users_data = {
    'id': [1, 2, 3],
    'name': ['Somchai Admin', 'Somsri Analyst', 'John Doe'],
    'email': ['admin@supplychain.com', 'analyst@supplychain.com', 'john@supplychain.com'],
    'password': ['admin1234', 'analyst1234', 'password123'],
    'role': ['Admin', 'Analyst', 'Analyst']
}
pd.DataFrame(users_data).to_csv('users.csv', index=False)

print("กำลังสร้าง category.csv...")
category = df[['Category Name']].drop_duplicates(subset=['Category Name']).reset_index(drop=True)
category.columns = ['category']
category.insert(0, 'id', range(1, 1 + len(category)))
category.to_csv('category.csv', index=False)

category_map = dict(zip(category['category'], category['id']))

# ---------------------------------------------------------
# 2. สร้าง products.csv (เหมือนเดิม)
# ---------------------------------------------------------
print("กำลังสร้าง products.csv...")
products = df[['Product Card Id', 'Product Name', 'Category Name', 'Product Price']].drop_duplicates()
products.columns = ['id', 'product_name', 'category_id', 'price']
products['category_id'] = products['category_id'].map(category_map)
products['stock'] = [random.randint(100, 1000) for _ in range(len(products))]
products.to_csv('products.csv', index=False)

# ---------------------------------------------------------
# 3. สร้าง orders.csv (✨ ปรับปรุงใหม่: ลบแถวซ้ำตาม Order Id ให้เหลือ 1 ออเดอร์ต่อ 1 บรรทัด)
# ---------------------------------------------------------
# ---------------------------------------------------------
# 3. สร้าง orders.csv (✨ ปรับปรุง: ดึงรหัสลูกค้าตัวจริงมาใส่แทน 999)
# ---------------------------------------------------------
print("กำลังสร้าง orders.csv...")
# 💡 เปลี่ยนจากดึงแค่ Order Id มาเป็นดึง 'Customer Id' จริงๆ จากไฟล์ DataCo มาด้วย
orders = df[['Order Id', 'order date (DateOrders)', 'Customer Id', 'Order Region']].drop_duplicates(subset=['Order Id'])
# เปลี่ยนชื่อคอลัมน์ให้ตรงกับที่เราออกแบบไว้ใน Supabase
orders.columns = ['id', 'order_date', 'customer_id', 'region'] 
orders.to_csv('orders.csv', index=False)
# ---------------------------------------------------------
# 4. สร้าง order_details.csv (✨ ตารางใหม่ที่คุณทัก: เก็บข้อมูลสินค้าในแต่ละออเดอร์)
# ---------------------------------------------------------
print("กำลังสร้าง order_details.csv...")
# ดึงข้อมูลการจับคู่ ออเดอร์ + สินค้า + ยอดขาย + กำไร ของชิ้นนั้นๆ
order_details = df[['Order Id', 'Product Card Id', 'Sales', 'Order Profit Per Order']]
order_details.columns = ['order_id', 'product_id', 'sales', 'profit']
# ใส่ ID รันตัวเลขให้ตาราง order_details เป็น Primary Key
order_details.insert(0, 'id', range(1, 1 + len(order_details)))
order_details.to_csv('order_details.csv', index=False)

# ---------------------------------------------------------
# 5. สร้าง shipments.csv (✨ ปรับปรุง: อิงตาม Order Id ที่เป็น Unique แล้ว)
# ---------------------------------------------------------
print("กำลังสร้าง shipments.csv...")
shipments = df[['Order Id', 'Shipping Mode', 'Days for shipping (real)', 'Delivery Status']].drop_duplicates(subset=['Order Id'])
shipments.columns = ['order_id', 'shipping_mode', 'delivery_days', 'status']
shipments.insert(0, 'id', range(1, 1 + len(shipments)))
shipments.to_csv('shipments.csv', index=False)

print("คลีนข้อมูลและแยกตามหลัก Relational Database เสร็จเรียบร้อย!")
